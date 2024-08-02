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
		File   string    `yaml:"file"`
		Token  EnvString `yaml:"token"`
		Expose []string  `yaml:"expose"`
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
		return ""
	}

	if s, ok := v.(string); ok {
		return EnvString(s).String()
	}

	return fmt.Sprint(v)
}

func (m M) GetOr(name, def string) string {
	v := m.Get(name)
	if v == "" {
		return def
	}
	return v
}

type EnvString string

func (e EnvString) IsEnv() bool {
	return e[0] == '$'
}

func (e EnvString) String() string {
	if e.IsEnv() {
		return os.Getenv(string(e[1:]))
	}
	return string(e)
}
