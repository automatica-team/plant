package plant

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	PrefixPlant   = "plant/"
	PrefixPrivate = "x/"
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

func (c Config) OfPrefix(ms []M, prefix string) (filtered []M) {
	for _, m := range ms {
		if strings.HasPrefix(m.Name(), prefix) {
			filtered = append(filtered, m)
		}
	}
	return filtered
}

func (c Config) V(m M) V {
	v := viper.New()
	v.MergeConfigMap(m)
	return V{Viper: v}
}

type (
	M map[string]any
	V struct{ *viper.Viper }
)

func (m M) Name() string {
	return m["import"].(string)
}

func (v V) GetEnv(name string) string {
	s := v.GetString(name)
	if s == "" {
		return s
	}
	return EnvString(s).String()
}

type EnvString string

func (e EnvString) IsEnv() bool {
	return len(e) > 0 && e[0] == '$'
}

func (e EnvString) String() string {
	if e.IsEnv() {
		return os.Getenv(string(e[1:]))
	}
	return string(e)
}
