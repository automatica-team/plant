package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	Version = "0.1"
	Header  = "🤖 Plant v" + Version
)

func main() {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.AddCommand(version)
	cmd.AddCommand(run)

	run.Flags().String("replace", "", "replace directive for go.mod")

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var (
	cmd = &cobra.Command{
		Use:   "plant",
		Short: Header,
		RunE:  func(c *cobra.Command, _ []string) error { return c.Help() },
	}
	version = &cobra.Command{
		Use:   "version",
		Short: "Print version and quit",
		Run:   func(_ *cobra.Command, _ []string) { fmt.Println(Header) },
	}
	run = &cobra.Command{
		Use:   "run [OPTIONS] PATH",
		Short: "Create and run a new bot from a config",
		RunE:  Run,
	}
)
