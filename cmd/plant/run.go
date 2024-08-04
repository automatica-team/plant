package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"automatica.team/plant"
	"automatica.team/plant/cmd/plant/exec"
	"automatica.team/plant/cmd/plant/template"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

type RunData struct {
	*plant.Plant
	ModName string
}

func Run(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return c.Usage()
	}

	project := args[0]
	exec.Cd(project)

	var (
		configPath = filepath.Join(project, "plant.yml")
		goModFile  = filepath.Join(project, "go.mod")
		botYmlFile = filepath.Join(project, "bot.yml")
		mainGoFile = filepath.Join(project, "main.gen.go")
		dotEnvFile = filepath.Join(project, ".env")
	)

	pl, err := plant.New(configPath)
	if err != nil {
		return fmt.Errorf("failed to parse plant config: %w", err)
	}

	defer func() {
		for _, path := range []string{mainGoFile, botYmlFile} {
			os.Remove(path)
		}
	}()

	run := RunData{
		Plant:   pl,
		ModName: project,
	}

	fmt.Println("[📝] Reading .env file")
	{
		if err := godotenv.Load(dotEnvFile); err != nil {
			return err
		}
	}

	fmt.Println("[📦] Creating go.mod file")
	{
		if _, err := os.Stat(goModFile); errors.Is(err, os.ErrNotExist) {
			if run.ModName == "." {
				wd, _ := os.Getwd()
				run.ModName = filepath.Base(wd)
			}

			if err := exec.Exec("go", "mod", "init", run.ModName); err != nil {
				return err
			}

			replace, _ := c.Flags().GetString("replace")
			if replace == "" {
				replace = "github.com/automatica-team/plant@latest"
			}

			replace = "automatica.team/plant=" + replace
			if err := exec.ExecSilent("go", "mod", "edit", "-replace", replace); err != nil {
				return err
			}
		} else {
			data, err := os.ReadFile(goModFile)
			if err != nil {
				return err
			}
			run.ModName = modfile.ModulePath(data)
		}
	}

	fmt.Println("[⚙️] Generating main.gen.go")
	{
		var buf bytes.Buffer
		if err := template.Run.ExecuteTemplate(&buf, "main.go", run); err != nil {
			return err
		}
		data, err := format.Source(buf.Bytes())
		if err != nil {
			return err
		}
		if err := os.WriteFile(mainGoFile, data, 0644); err != nil {
			return err
		}
	}

	fmt.Println("[🤖] Generating bot.yml file")
	{
		var buf bytes.Buffer
		if err := template.Run.ExecuteTemplate(&buf, "bot.yml", run); err != nil {
			return err
		}
		if err := os.WriteFile(botYmlFile, buf.Bytes(), 0644); err != nil {
			return err
		}
	}

	fmt.Println("[📥] Downloading Go modules")
	{
		if err := exec.Exec("go", "mod", "tidy"); err != nil {
			return err
		}
	}

	fmt.Println("[🚀] Running the bot")
	{
		return exec.Exec("go", "run", filepath.Base(mainGoFile))
	}
}
