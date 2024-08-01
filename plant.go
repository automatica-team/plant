package plant

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

type (
	Config struct {
		Bot  ConfigBot `yaml:"bot"`
		Deps []M       `yaml:"deps"`
		Mods []M       `yaml:"mods"`
	}

	ConfigBot struct {
		Expose []string `yaml:"expose"`
	}
)

type Plant struct {
	Config
	mods map[string]Mod
}

func New(path string) (*Plant, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &Plant{
		Config: conf,
		mods:   make(map[string]Mod),
	}, nil
}

func (p *Plant) Build(b *Bot) error {
	if len(p.Mods) == 0 {
		return errors.New("plant: no mods to build")
	}

	for _, m := range p.Mods {
		mod := p.mods[m.Name()]
		if err := mod.Import(m); err != nil {
			return err
		}
	}

	for _, end := range p.Bot.Expose {
		b.Handle(end, func(c tele.Context) error {
			for _, mod := range p.mods {
				if h := mod.Handler(end); h != nil {
					if err := h(c); err != nil {
						return err
					}
				}
			}
			return nil
		})
	}

	return nil
}

type M map[string]any

func (m M) Name() string {
	return m["import"].(string)
}

func (m M) Get(name string) string {
	v, ok := m[name]
	if !ok {
		panic(name)
	}

	if s, ok := v.(string); ok {
		if s[0] == '$' {
			return os.Getenv(s[1:])
		}
	}

	return fmt.Sprint(v)
}
