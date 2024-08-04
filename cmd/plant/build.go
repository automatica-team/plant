package main

import (
	"fmt"
	"path/filepath"

	"automatica.team/plant"
	"automatica.team/plant/cmd/plant/do"
	"automatica.team/plant/cmd/plant/exec"

	"github.com/spf13/cobra"
)

func Build(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return c.Usage()
	}

	project := args[0]
	exec.Cd(project)

	var (
		path   = filepath.Join(project, "plant.yml")
		tag, _ = c.Flags().GetString("tag")
	)

	pl, err := plant.New(path)
	if err != nil {
		return fmt.Errorf("failed to parse plant config: %w", err)
	}

	doData := do.Ctx{
		Plant:   pl,
		Project: project,
		ModName: project,
	}

	purge, err := do.Base(doData)
	defer purge() // purge always before err check
	if err != nil {
		return err
	}

	remove, err := do.Dockerfile(project)
	if err != nil {
		return err
	}
	defer remove()

	fmt.Println("[🐳] Building the Docker image")
	if err := exec.Run("docker", buildDockerBuildArgs(tag)...); err != nil {
		return err
	}

	if tag != "" {
		fmt.Println("[🐳] Pushing the Docker image")
		return exec.Run("docker", "push", tag)
	}

	return nil
}

func buildDockerBuildArgs(tag string) []string {
	dockerBuildArgs := []string{"build"}
	if tag != "" {
		dockerBuildArgs = append(dockerBuildArgs, "-t", tag)
	}
	return append(dockerBuildArgs, ".")
}
