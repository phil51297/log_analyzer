package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "log_analyzer",
	Short: "Un outil d'analyse de logs en parallèle",
	Long:  `loganalyzer est un outil CLI conçu pour analyser des fichiers de logs`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Erreur affichée sur la sortie d'erreur standard (stderr)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		// Le programme se termine avec un code d'erreur non nul (1a0 pour indiquer un echec
		os.Exit(1)
	}
}
