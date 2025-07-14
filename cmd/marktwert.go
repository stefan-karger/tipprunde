package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// marktwertCmd represents the marktwert command
var marktwertCmd = &cobra.Command{
	Use:   "marktwert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
goto quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("marktwert called")
	},
}

func init() {
	rootCmd.AddCommand(marktwertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// marktwertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// marktwertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
