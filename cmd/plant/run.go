package main

import (
	"fmt"
	"path/filepath"

	"automatica.team/plant"
	"automatica.team/plant/cmd/plant/do"
	"automatica.team/plant/cmd/plant/exec"

	"github.com/spf13/cobra"
)

func Run(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return c.Usage()
	}

	project := args[0]
	exec.Cd(project)

	var (
		path       = filepath.Join(project, "plant.yml")
		replace, _ = c.Flags().GetString("replace")
	)

	pl, err := plant.New(path)
	if err != nil {
		return fmt.Errorf("failed to parse plant config: %w", err)
	}

	if err := do.DotEnv(project); err != nil {
		return err
	}

	ctx := do.Ctx{
		Plant:   pl,
		Project: project,
		ModName: project,
		Replace: replace,
	}

	purge, err := do.Base(ctx)
	defer purge() // always before err check
	if err != nil {
		return err
	}

	fmt.Println("[🚀] Running the bot")
	return exec.Run("go", "run", filepath.Base("main.go"))
}
