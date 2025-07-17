package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tipprunde",
	Short: "Kurzbeschreibung",
	Long:  `Lange Beschreibung.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(marktwertCmd)
	rootCmd.AddCommand(ergebnisseCmd)
}
