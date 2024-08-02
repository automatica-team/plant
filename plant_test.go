package plant

import (
	"testing"

	"github.com/stretchr/testify/require"
	tele "gopkg.in/telebot.v3"
)

func TestPlant(t *testing.T) {
	p, err := New("plant.yml")
	require.NoError(t, err)
	require.NotEmpty(t, p.Config)

	dep := &TestDep{}
	mod := &TestMod{}
	p.Inject(dep)
	p.Add(mod)

	bot, err := p.Build()
	require.NoError(t, err)
	require.NotNil(t, bot)

	// Check if the deps and mods are filled
	require.Equal(t, dep, p.deps["plant/test"])
	require.Equal(t, mod, p.mods["plant/test"])

	// Check if the mod fields were injected
	require.Equal(t, bot, mod.Handler.b)
	require.Equal(t, bot, mod.b)
	require.Equal(t, dep, mod.td)

	// Check if the handlers were added
	require.Len(t, bot.h["/start"], 1)
}

type TestDep struct{}

func (TestDep) Name() string {
	return "plant/test"
}

func (t TestDep) Import(_ M) error {
	return nil
}

type TestMod struct {
	Handler
	b  *Bot     `plant:"bot"`
	td *TestDep `plant:"dep:plant/test"`
}

func (TestMod) Name() string {
	return "plant/test"
}

func (mod TestMod) Import(_ M) error {
	mod.Handle("/start", func(c tele.Context) error {
		return nil
	})

	return nil
}
