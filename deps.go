package plant

import (
	"errors"
	"fmt"
)

type Deps map[string]Dep

func Inject(dep Dep) {
	globalDeps[dep.Name()] = dep
}

type Dep interface {
	Name() string
	Import(M) error
}

var globalDeps = make(Deps)

func (p *Plant) Connect() (Deps, error) {
	deps := make(Deps, len(p.Deps))
	for _, m := range p.Deps {
		dep, ok := globalDeps[m.Name()]
		if !ok {
			return nil, errors.New("plant: dependency not injected")
		}
		if err := dep.Import(m); err != nil {
			return nil, fmt.Errorf("plant: dependency connect: %w", err)
		}
		deps[dep.Name()] = dep
	}
	return deps, nil
}
