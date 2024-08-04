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
	cmd.AddCommand(build)

	run.Flags().String("replace", "", "replace directive for go.mod")
	build.Flags().StringP("tag", "t", "", "tag for the Docker image")

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
		Use:     "run [OPTIONS] PATH",
		Short:   "Create and run a new bot from a config",
		Example: "plant run demo",
		RunE:    Run,
	}
	build = &cobra.Command{
		Use:     "build [OPTIONS] PATH",
		Short:   "Builds a Docker image for the bot",
		Example: "plant build -t ghcr.io/automatica-team/bots/demo demo",
		RunE:    Build,
	}
)
