package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"time"

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
		mainGoFile = filepath.Join(project, "main.gen.go")
		botYmlFile = filepath.Join(project, "bot.yml")
		goModFile  = filepath.Join(project, "go.mod")
		goSumFile  = filepath.Join(project, "go.sum")
		dotEnvFile = filepath.Join(project, ".env")
	)

	pl, err := plant.New(path)
	if err != nil {
		return fmt.Errorf("failed to parse plant config: %w", err)
	}

	defer func() {
		for _, path := range []string{mainGoFile, botYmlFile, goModFile, goSumFile} {
			os.Remove(path)
		}
	}()

	fmt.Println("📝 Reading .env file")
	{
		if err := godotenv.Load(dotEnvFile); err != nil {
			return err
		}
	}

	fmt.Println("⚙️ Generating main.gen.go")
	{
		var buf bytes.Buffer
		if err := template.Run.ExecuteTemplate(&buf, "main.go", pl); err != nil {
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

	fmt.Println("🤖 Generating bot.yml file")
	{
		var buf bytes.Buffer
		if err := template.Run.ExecuteTemplate(&buf, "bot.yml", pl); err != nil {
			return err
		}
		if err := os.WriteFile(botYmlFile, buf.Bytes(), 0644); err != nil {
			return err
		}
	}

	fmt.Println("📦 Generating go.mod file")
	{
		if _, err := os.Stat(goModFile); errors.Is(err, os.ErrNotExist) {
			modName := project
			if modName == "." {
				wd, _ := os.Getwd()
				modName = filepath.Base(wd)
			}

			if err := exec.Exec("go", "mod", "init", modName); err != nil {
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
		}
	}

	fmt.Println("📥 Downloading Go modules")
	{
		if err := exec.Exec("go", "mod", "tidy"); err != nil {
			return err
		}
	}

	fmt.Println("🚀 Running the bot")
	{
		cmd, err := exec.Command("go", "run", filepath.Base(mainGoFile))
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
		fmt.Print("Press any key to exit...")
		fmt.Scanln()

		return cmd.Process.Kill()
	}
}
