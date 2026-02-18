package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	// Global flags set from root command.
	JSONOutput bool
	NoColor    bool
	Quiet      bool
)

// Init initializes output settings based on global flags and environment.
func Init() {
	if NoColor || os.Getenv("NO_COLOR") != "" {
		color.NoColor = true
	}
}

// Success prints a success message to stderr.
func Success(msg string) {
	if Quiet {
		return
	}
	if JSONOutput {
		printJSON(map[string]string{"status": "ok", "message": msg})
		return
	}
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Fprintf(os.Stderr, "%s %s\n", green("✓"), msg)
}

// Error prints an error message with optional help text to stderr.
func Error(problem, help string) {
	if Quiet {
		return
	}
	if JSONOutput {
		m := map[string]string{"error": problem}
		if help != "" {
			m["help"] = help
		}
		printJSON(m)
		return
	}
	red := color.New(color.FgRed).SprintFunc()
	fmt.Fprintf(os.Stderr, "%s Error: %s\n", red("✗"), problem)
	if help != "" {
		fmt.Fprintf(os.Stderr, "  Hint: %s\n", help)
	}
}

// Result prints a result to stdout.
func Result(text string) {
	fmt.Print(text)
}

// ResultLn prints a result with newline to stdout.
func ResultLn(text string) {
	fmt.Println(text)
}

func printJSON(v interface{}) {
	enc := json.NewEncoder(os.Stderr)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(v)
}

// PrintJSON prints a value as JSON to stdout.
func PrintJSON(v interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	_ = enc.Encode(v)
}
