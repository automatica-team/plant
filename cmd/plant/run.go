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
)

func Run(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return c.Usage()
	}

	project := args[0]
	exec.Cd(project)

	var (
		path       = filepath.Join(project, "plant.yml")
		mainFile   = filepath.Join(project, "main.gen.go")
		goModFile  = filepath.Join(project, "go.mod")
		dotEnvFile = filepath.Join(project, ".env")
	)

	pl, err := plant.New(path)
	if err != nil {
		return fmt.Errorf("failed to parse plant config: %w", err)
	}

	// 1.
	fmt.Println("⚙️ Generating main.gen.go")

	var buf bytes.Buffer
	if err := template.Run.Execute(&buf, pl); err != nil {
		return err
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	if err := os.WriteFile(mainFile, data, 0644); err != nil {
		return err
	}

	// 2.
	fmt.Println("📦 Generating go.mod file")
	if _, err := os.Stat(goModFile); errors.Is(err, os.ErrNotExist) {
		if err := exec.CommandSilent("go", "mod", "init", project); err != nil {
			return err
		}
		if err := exec.CommandSilent("go", runGoModEditReplace...); err != nil {
			return err
		}
	}

	// 3.
	fmt.Println("📥 Downloading Go modules")
	if err := exec.Command("go", "mod", "tidy"); err != nil {
		return err
	}

	// 4.
	fmt.Println("📝 Reading .env file")
	if err := godotenv.Load(dotEnvFile); err != nil {
		return err
	}

	// 4.
	fmt.Println("🚀 Running the bot")
	if err := exec.Command("go", "run", filepath.Base(mainFile)); err != nil {
		return err
	}

	return nil
}

var runGoModEditReplace = []string{
	"mod", "edit", "-replace",
	"automatica.team/plant=github.com/automatica-team/plant@latest",
}
