package cmd

import (
	"os"

	"github.com/dontbeidle/kaalin/internal/output"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kaalin",
	Short: "A CLI tool for the Karakalpak language",
	Long: `kaalin — a text conversion tool for the Karakalpak language.

Features:
  - Latin ↔ Cyrillic script conversion
  - Number → words conversion
  - Upper / Lower case (Karakalpak alphabet aware)`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		output.Init()
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&output.JSONOutput, "json", "j", false, "Output as JSON")
	rootCmd.PersistentFlags().BoolVar(&output.NoColor, "no-color", false, "Disable colored output")
	rootCmd.PersistentFlags().BoolVarP(&output.Quiet, "quiet", "q", false, "Only print the result")

	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(numberCmd)
	rootCmd.AddCommand(caseCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)
}
