package plant

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Bot  ConfigBot `yaml:"bot"`
		Deps []M       `yaml:"deps"`
		Mods []M       `yaml:"mods"`
	}

	ConfigBot struct {
		File   string   `yaml:"file"`
		Expose []string `yaml:"expose"`
	}
)

func Parse(path string) (c Config, _ error) {
	v := viper.New()
	v.SetConfigFile(path)

	// Defaults
	v.SetDefault("bot.file", "bot.yml")

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	return c, v.Unmarshal(&c)
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
