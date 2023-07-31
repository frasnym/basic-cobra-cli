// cmd/subcommand.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "subcommand",
	Short: "A brief description of the subcommand",
	Long: `A longer description that explains the subcommand
in detail. For example, you can mention how it complements
the main command and what options it supports.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This function will be executed when the "subcommand" is called
		fmt.Println("Running the subcommand!")
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
