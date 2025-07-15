package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var marktwertCmd = &cobra.Command{
	Use:   "marktwert",
	Short: "Kurzbeschreibung",
	Long: `Lange Beschreibung.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("marktwert called")
		content := GetURLContent("https://www.transfermarkt.de/fc-bayern-munchen/kader/verein/27/saison_id/2025/plus/1")
		fmt.Println(content)
	},
}

func init() {
	rootCmd.AddCommand(marktwertCmd)
}
