package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/dontbeidle/kaalin/internal/number"
	"github.com/dontbeidle/kaalin/internal/output"
	"github.com/spf13/cobra"
)

var numberCmd = &cobra.Command{
	Use:     "number <son>",
	Short:   "Convert a number to words",
	Aliases: []string{"n"},
	Long: `Convert a number to Karakalpak words.

For negative numbers use --: kaalin number -- -5`,
	// Disable flag parsing to allow negative numbers like -5.
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Manually parse flags
		var cyr bool
		var cleanArgs []string
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "--cyr", "-c":
				cyr = true
			case "--json", "-j":
				output.JSONOutput = true
			case "--no-color":
				output.NoColor = true
			case "--quiet", "-q":
				output.Quiet = true
			case "--help", "-h":
				return cmd.Help()
			case "--":
				// Everything after -- is arguments
				cleanArgs = append(cleanArgs, args[i+1:]...)
				i = len(args) // break loop
			default:
				cleanArgs = append(cleanArgs, args[i])
			}
		}

		output.Init()

		input, err := getNumberInput(cleanArgs)
		if err != nil {
			output.Error(err.Error(), "")
			os.Exit(2)
		}

		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			output.Error(
				fmt.Sprintf("\"%s\" is not a valid number", input),
				"enter an integer or decimal (e.g. 123, 12.75, -5)",
			)
			os.Exit(2)
		}

		script := "lat"
		if cyr {
			script = "cyr"
		}

		result, err := number.ToWord(num, script)
		if err != nil {
			output.Error(err.Error(), "")
			os.Exit(1)
		}

		if output.JSONOutput {
			output.PrintJSON(map[string]string{"result": result})
		} else {
			output.ResultLn(result)
		}

		return nil
	},
}

func getNumberInput(args []string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read stdin: %s", err)
		}
		return strings.TrimSpace(string(data)), nil
	}

	return "", fmt.Errorf("no number provided")
}
