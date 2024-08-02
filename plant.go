package plant

import (
	"fmt"
)

type Plant struct {
	Config

	deps map[string]Dep
	mods map[string]Mod
}

func New(path string) (*Plant, error) {
	conf, err := Parse(path)
	if err != nil {
		return nil, err
	}

	return &Plant{
		Config: conf,
		deps:   make(map[string]Dep),
		mods:   make(map[string]Mod),
	}, nil
}

func (p *Plant) Build() (*Bot, error) {
	b, err := p.composeBot()
	if err != nil {
		return nil, fmt.Errorf("plant: %w", err)
	}

	if err := p.injectMods(b); err != nil {
		return nil, fmt.Errorf("plant: %w", err)
	}
	if err := p.importDeps(); err != nil {
		return nil, fmt.Errorf("plant: %w", err)
	}
	if err := p.importMods(); err != nil {
		return nil, fmt.Errorf("plant: %w", err)
	}

	return b, p.exposeBot(b)
}
