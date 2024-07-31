package plant

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

type Config struct {
	Bot struct {
		Config string `yaml:"config"`
	} `yaml:"bot"`

	Joints []string `yaml:"joints"`
	Deps   []M      `yaml:"deps"`
	Parts  []M      `yaml:"parts"`
}

type Plant struct {
	Config
	parts map[string]Part
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
		parts:  make(map[string]Part),
	}, nil
}

func (p *Plant) Build(b *Bot) error {
	if len(p.Parts) == 0 {
		return errors.New("plant: no parts to build")
	}

	for _, m := range p.Parts {
		part := p.parts[m.Name()]
		if err := part.Import(m); err != nil {
			return err
		}
	}

	for _, joint := range p.Joints {
		b.Handle(joint, func(c tele.Context) error {
			for _, part := range p.parts {
				if h := part.Handler(joint); h != nil {
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
