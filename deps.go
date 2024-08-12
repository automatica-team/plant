package plant

import (
	"fmt"
)

type Dep interface {
	Name() string
	Import(V) error
}

func (p *Plant) Inject(d Dep) {
	p.deps[d.Name()] = d
}

func (p *Plant) importDeps() error {
	for _, m := range p.Deps {
		name := m.Name()

		dep, ok := p.deps[name]
		if !ok {
			return fmt.Errorf("importDeps: dep %s not injected", name)
		}

		if err := dep.Import(p.Config.V(m)); err != nil {
			return fmt.Errorf("importDeps: (%s) %w", name, err)
		}
	}

	return nil
}
