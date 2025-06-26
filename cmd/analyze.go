package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/phil51297/log_analyzer/internal/analyzer"
	"github.com/phil51297/log_analyzer/internal/config"
	"github.com/phil51297/log_analyzer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configFile string
	outputFile string
	verbose    bool
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyser des fichiers de logs en parallèle",
	Long: `Analyser des fichiers de logs configurés dans un fichier JSON.
	
Le fichier de configuration doit contenir une liste de logs avec leurs identifiants, chemins et types.
L'analyse est effectuée en parallèle pour tous les fichiers spécifiés.`,
	Run: analyzeRun,
}

func analyzeRun(cmd *cobra.Command, args []string) {
	if configFile == "" {
		fmt.Println("❌ Erreur: Le fichier de configuration est requis.")
		fmt.Println("Utilisation: log_analyzer analyze --config config.json")
		os.Exit(1)
	}

	analyzeWithConfig()
}

func analyzeWithConfig() {
	fmt.Printf("📋 Chargement de la configuration: %s\n", configFile)

	configs, err := config.LoadConfig(configFile)
	if err != nil {
		fmt.Printf("❌ Erreur lors du chargement de la configuration: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📊 Démarrage de l'analyse de %d fichier(s) de log en parallèle...\n\n", len(configs))

	resultsChan := make(chan analyzer.AnalysisResult, len(configs))

	var wg sync.WaitGroup

	for i, logConfig := range configs {
		wg.Add(1)
		go func(index int, config config.LogConfig) {
			defer wg.Done()

			if verbose {
				fmt.Printf("  [%d/%d] 🚀 Démarrage de l'analyse de %s (%s)\n", index+1, len(configs), config.ID, config.Path)
			}

			result := analyzer.AnalyzeLog(config)

			resultsChan <- result

			if verbose {
				status := "✅"
				if result.Status == "FAILED" {
					status = "❌"
				}
				fmt.Printf("  [%d/%d] %s Terminé: %s\n", index+1, len(configs), status, config.ID)
			}
		}(i, logConfig)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []analyzer.AnalysisResult
	successCount := 0

	for result := range resultsChan {
		results = append(results, result)
		if result.Status == "OK" {
			successCount++
		}
	}

	fmt.Printf("\n✅ Analyse terminée: %d/%d fichiers traités avec succès\n\n", successCount, len(configs))

	displayResults(results)

	if outputFile != "" {
		exportResults(results)
	}
}

func displayResults(results []analyzer.AnalysisResult) {
	fmt.Println("📊 RÉSULTATS D'ANALYSE:")
	fmt.Println("=" + fmt.Sprintf("%*s", 80, "="))

	for _, result := range results {
		fmt.Printf("\n🆔 ID: %s\n", result.LogID)
		fmt.Printf("📄 Chemin: %s\n", result.FilePath)

		if result.Status == "OK" {
			fmt.Printf("✅ Statut: %s\n", result.Status)
			fmt.Printf("💬 Message: %s\n", result.Message)
		} else {
			fmt.Printf("❌ Statut: %s\n", result.Status)
			fmt.Printf("💬 Message: %s\n", result.Message)
			if result.ErrorDetails != "" {
				fmt.Printf("🔍 Détails de l'erreur: %s\n", result.ErrorDetails)
			}
		}
		fmt.Println("-" + fmt.Sprintf("%*s", 50, "-"))
	}
}

func exportResults(results []analyzer.AnalysisResult) {
	fmt.Printf("\n💾 Export des résultats vers: %s\n", outputFile)

	err := reporter.ExportToJSON(results, outputFile)
	if err != nil {
		fmt.Printf("❌ Erreur lors de l'export: %v\n", err)
		return
	}

	fmt.Printf("✅ Résultats exportés avec succès dans %s\n", outputFile)
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configFile, "config", "c", "", "Fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Fichier de sortie pour exporter les résultats en JSON")
	analyzeCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Mode verbeux pour afficher les détails")

	analyzeCmd.MarkFlagRequired("config")
}
