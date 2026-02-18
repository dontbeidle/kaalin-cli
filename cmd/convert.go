package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dontbeidle/kaalin/internal/converter"
	"github.com/dontbeidle/kaalin/internal/output"
	"github.com/spf13/cobra"
)

var (
	toCyr   bool
	toLat   bool
	file    string
	outFile string
	inPlace bool
)

var convertCmd = &cobra.Command{
	Use:     "convert [text]",
	Short:   "Convert between Latin and Cyrillic scripts",
	Aliases: []string{"c"},
	Long: `Convert text between Latin and Cyrillic scripts.

Input sources (by priority):
  1. Argument: kaalin convert "Sálem"
  2. Pipe:     echo "Sálem" | kaalin convert
  3. File:     kaalin convert -f input.txt
  4. Interactive mode (if none of the above)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if toCyr && toLat {
			output.Error("--to-cyr and --to-lat cannot be used together", "choose only one")
			os.Exit(2)
		}

		if inPlace && file == "" {
			output.Error("--in-place requires --file (-f)", "add the --file flag")
			os.Exit(2)
		}

		// Get input text
		text, err := getInput(args)
		if err != nil {
			output.Error(err.Error(), "")
			os.Exit(1)
		}

		// Detect direction if not specified
		if !toCyr && !toLat {
			script := converter.DetectScript(text)
			if script == "cyrillic" {
				toLat = true
			} else {
				toCyr = true
			}
		}

		// Convert
		var result string
		if toCyr {
			result = converter.Latin2Cyrillic(text)
		} else {
			result = converter.Cyrillic2Latin(text)
		}

		// Output
		if outFile != "" {
			if err := os.WriteFile(outFile, []byte(result), 0644); err != nil {
				output.Error(fmt.Sprintf("failed to write file: %s", err), "")
				os.Exit(1)
			}
			output.Success(fmt.Sprintf("Converted: %s", outFile))
			return nil
		}

		if inPlace {
			if err := os.WriteFile(file, []byte(result), 0644); err != nil {
				output.Error(fmt.Sprintf("failed to update file: %s", err), "")
				os.Exit(1)
			}
			output.Success(fmt.Sprintf("File updated: %s", file))
			return nil
		}

		if output.JSONOutput {
			output.PrintJSON(map[string]string{"result": result})
		} else {
			output.ResultLn(result)
		}

		return nil
	},
}

func getInput(args []string) (string, error) {
	// Priority 1: inline argument
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// Priority 2: stdin/pipe
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("failed to read stdin: %s", err)
		}
		return strings.TrimRight(string(data), "\n"), nil
	}

	// Priority 3: file
	if file != "" {
		data, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("failed to read file: %s", err)
		}
		return string(data), nil
	}

	// Priority 4: interactive mode
	return interactiveInput()
}

func interactiveInput() (string, error) {
	direction := "Cyrillic"
	if toLat {
		direction = "Latin"
	}
	fmt.Fprintf(os.Stderr, "Enter text (converting to %s, Ctrl+D to finish):\n", direction)
	fmt.Fprint(os.Stderr, "> ")

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		fmt.Fprint(os.Stderr, "> ")
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("no text provided")
	}

	return strings.Join(lines, "\n"), nil
}

func init() {
	convertCmd.Flags().BoolVarP(&toCyr, "to-cyr", "c", false, "Convert to Cyrillic")
	convertCmd.Flags().BoolVarP(&toLat, "to-lat", "l", false, "Convert to Latin")
	convertCmd.Flags().StringVarP(&file, "file", "f", "", "Input file")
	convertCmd.Flags().StringVarP(&outFile, "output", "o", "", "Output file")
	convertCmd.Flags().BoolVar(&inPlace, "in-place", false, "Edit file in place (requires --file)")
}
