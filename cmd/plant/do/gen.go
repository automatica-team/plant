package do

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"automatica.team/plant/cmd/plant/exec"
	"automatica.team/plant/cmd/plant/tmpl"

	"golang.org/x/mod/modfile"
)

const (
	goModReplace = "github.com/automatica-team/plant@latest"
)

func GoMod(ctx Ctx) (string, error) {
	if ctx.Replace == "" {
		ctx.Replace = goModReplace
	}

	var (
		path    = filepath.Join(ctx.Project, "go.mod")
		modName = ctx.Project
		replace = "automatica.team/plant=" + ctx.Replace
	)

	fmt.Println("[📦] Creating go.mod file")
	{
		if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
			data, err := os.ReadFile(path)
			if err != nil {
				return modName, err
			}
			return modfile.ModulePath(data), nil
		}

		if modName == "." {
			wd, _ := os.Getwd()
			modName = filepath.Base(wd)
		}

		if err := exec.Run("go", "mod", "init", modName); err != nil {
			return modName, err
		}

		if err := exec.RunSilent("go", "mod", "edit", "-replace", replace); err != nil {
			return modName, err
		}
	}

	return modName, nil
}

func GoSum() error {
	fmt.Println("[📥] Downloading Go modules")
	{
		return exec.Run("go", "mod", "tidy")
	}
}

func MainGo(ctx Ctx) (func(), error) {
	path := filepath.Join(ctx.Project, "main.go")

	fmt.Println("[⚙️] Generating main.go")
	{
		var buf bytes.Buffer
		if err := tmpl.Run.ExecuteTemplate(&buf, "main.go", ctx); err != nil {
			return nil, err
		}
		data, err := format.Source(buf.Bytes())
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, data, 0644); err != nil {
			return nil, err
		}
	}

	return func() { os.Remove(path) }, nil
}

func BotYml(ctx Ctx) (func(), error) {
	path := filepath.Join(ctx.Project, "bot.yml")

	fmt.Println("[⚙️] Generating bot.yml")
	{
		var buf bytes.Buffer
		if err := tmpl.Run.ExecuteTemplate(&buf, "bot.yml", ctx); err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
			return nil, err
		}
	}

	return func() { os.Remove(path) }, nil
}
