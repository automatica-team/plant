package main

import (
	"automatica.team/plant"
	"automatica.team/plant/cmd/plant/do"
	"automatica.team/plant/cmd/plant/exec"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

func Gen(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return c.Usage()
	}

	project := args[0]
	exec.Cd(project)

	var (
		path = filepath.Join(project, "plant.yml")
	)

	pl, err := plant.New(path)
	if err != nil {
		return fmt.Errorf("failed to parse plant config: %w", err)
	}

	ctx := do.Ctx{
		Plant:   pl,
		Project: project,
		ModName: project,
	}

	_, err = do.Base(ctx)

	if err != nil {
		return err
	}

	return nil
}
