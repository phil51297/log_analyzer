package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "log_analyzer",
	Short: "Un outil d'analyse de logs en parallèle",
	Long: `loganalyzer est un outil CLI conçu pour analyser des fichiers de logs en parallèle.
	
Utilise des goroutines pour traiter plusieurs fichiers de logs simultanément 
et génère des rapports détaillés avec gestion d'erreurs robuste.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
