package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const Header = "🤖 Plant " + verPlant

func main() {
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.AddCommand(version)
	cmd.AddCommand(run)
	cmd.AddCommand(build)
	cmd.AddCommand(gen)

	run.Flags().String("replace", "", "replace directive for go.mod")
	build.Flags().StringP("tag", "t", "", "tag for the Docker image")
	build.Flags().StringP("platform", "p", "", "platform for the Docker image")

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
		RunE:  Version,
	}
	run = &cobra.Command{
		Use:     "run [OPTIONS] PATH",
		Short:   "Create and run a new bot from a config",
		Example: "plant run demo",
		RunE:    Run,
	}
	gen = &cobra.Command{
		Use:     "gen [OPTIONS] PATH",
		Short:   "Create and generate boilerplate code",
		Example: "plant gen demo",
		RunE:    Gen,
	}
	build = &cobra.Command{
		Use:     "build [OPTIONS] PATH",
		Short:   "Builds a Docker image for the bot",
		Example: "plant build -t ghcr.io/automatica-team/bots/demo demo",
		RunE:    Build,
	}
)
