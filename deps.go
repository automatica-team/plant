package plant

import (
	"errors"
	"fmt"
	"os"
)

type Deps map[string]Dep

func Inject(dep Dep) {
	globalDeps[dep.Name()] = dep
}

type Dep interface {
	Name() string
	Connect(M) error
}

var globalDeps = make(Deps)

func (p *Plant) Connect() (Deps, error) {
	deps := make(Deps, len(p.Deps))
	for name, conf := range p.Deps {
		dep, ok := globalDeps[name]
		if !ok {
			return nil, errors.New("plant: dependency not injected")
		}
		if err := dep.Connect(conf); err != nil {
			return nil, fmt.Errorf("plant: dependency connect: %w", err)
		}
		deps[name] = dep
	}
	return deps, nil
}

type M map[string]any

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
