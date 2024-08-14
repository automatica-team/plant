package do

import (
	"fmt"
	"path/filepath"

	"automatica.team/plant"
	"automatica.team/plant/cmd/constants"
	"github.com/joho/godotenv"
)

var emptyFunc = func() {}

type Ctx struct {
	*plant.Plant
	Project string
	ModName string
	Replace string
}

func Base(ctx Ctx) (func(), error) { // nil-safe
	modName, err := GoMod(ctx)
	if err != nil {
		return emptyFunc, err
	}

	ctx.ModName = modName

	err = Get("gopkg.in/telebot.v3@" + constants.VerTelebot)
	if err != nil {
		return nil, err
	}

	remove1, err := MainGo(ctx)
	if err != nil {
		return emptyFunc, err
	}

	remove2, err := BotYml(ctx)
	if err != nil {
		return remove1, err
	}

	purge := func() {
		remove1()
		remove2()
	}

	return purge, GoSum()
}

func DotEnv(project string) error {
	path := filepath.Join(project, ".env")

	fmt.Println("[📝] Reading .env file")
	{
		return godotenv.Load(path)
	}
}
