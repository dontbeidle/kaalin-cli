package cmd

import (
	"fmt"
	"runtime"

	"github.com/dontbeidle/kaalin/internal/output"
	"github.com/spf13/cobra"
)

// Set via ldflags at build time.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show version info",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		if output.JSONOutput {
			output.PrintJSON(map[string]string{
				"version": version,
				"commit":  commit,
				"date":    date,
				"go":      runtime.Version(),
				"os":      runtime.GOOS,
				"arch":    runtime.GOARCH,
			})
			return
		}

		fmt.Printf("kaalin %s (%s)\n", version, commit)
		fmt.Printf("  built:   %s\n", date)
		fmt.Printf("  go:      %s\n", runtime.Version())
		fmt.Printf("  os/arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}
