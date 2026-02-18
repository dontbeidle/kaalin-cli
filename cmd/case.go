package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dontbeidle/kaalin/internal/output"
	"github.com/dontbeidle/kaalin/internal/strutil"
	"github.com/spf13/cobra"
)

var caseCmd = &cobra.Command{
	Use:   "case",
	Short: "Change letter casing",
	Long:  `Convert text to upper or lower case (Karakalpak alphabet aware).`,
}

var upperCmd = &cobra.Command{
	Use:   "upper [text]",
	Short: "Convert to uppercase",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getCaseInput(args)
		if err != nil {
			output.Error(err.Error(), "")
			os.Exit(2)
		}

		result := strutil.Upper(text)

		if output.JSONOutput {
			output.PrintJSON(map[string]string{"result": result})
		} else {
			output.ResultLn(result)
		}
		return nil
	},
}

var lowerCmd = &cobra.Command{
	Use:   "lower [text]",
	Short: "Convert to lowercase",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getCaseInput(args)
		if err != nil {
			output.Error(err.Error(), "")
			os.Exit(2)
		}

		result := strutil.Lower(text)

		if output.JSONOutput {
			output.PrintJSON(map[string]string{"result": result})
		} else {
			output.ResultLn(result)
		}
		return nil
	},
}

func getCaseInput(args []string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read stdin: %s", err)
		}
		return strings.TrimRight(string(data), "\n"), nil
	}

	return "", fmt.Errorf("no text provided")
}

func init() {
	caseCmd.AddCommand(upperCmd)
	caseCmd.AddCommand(lowerCmd)
}
