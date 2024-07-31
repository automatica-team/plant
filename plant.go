package plant

import (
	"errors"

	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

type Config struct {
	Bot struct {
		Config string `yaml:"config"`
	} `yaml:"bot"`

	Deps   map[string]M `yaml:"deps"`
	Joints []string     `yaml:"joints"`
	Parts  []string     `yaml:"parts"`
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

	parts := make([]Part, 0, len(p.Parts))
	for _, name := range p.Parts {
		parts = append(parts, p.parts[name])
	}

	for _, part := range parts {
		if err := part.Prepare(); err != nil {
			return err
		}
	}

	for _, joint := range p.Joints {
		b.Handle(joint, func(c tele.Context) error {
			for _, part := range parts {
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
