package plant

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type Mod interface {
	Name() string
	Expose() []any
	Import(V) error
}

func (p *Plant) Add(mod Mod) {
	p.mods[mod.Name()] = mod
}

func (p *Plant) injectMods(b *Bot) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("injectMods: %v", r)
		}
	}()

	for _, mod := range p.mods {
		v := reflect.ValueOf(mod).Elem()

		for i := 0; i < v.NumField(); i++ {
			var (
				f   = v.Field(i)
				t   = f.Type()
				tag = v.Type().Field(i).Tag.Get("plant")
			)

			uf := reflect.NewAt(t, unsafe.Pointer(f.UnsafeAddr())).Elem()

			switch {
			case tag == "" && t.String() == "plant.Handler":
				uf.Set(reflect.ValueOf(Handler{b: b}))
			case tag == "bot":
				uf.Set(reflect.ValueOf(b))
			case strings.HasPrefix(tag, "dep:"):
				name := strings.TrimPrefix(tag, "dep:")

				dep, ok := p.deps[name]
				if !ok {
					return fmt.Errorf("injectMods: (%s) dep not injected", name)
				}

				uf.Set(reflect.ValueOf(dep))
			}
		}
	}

	return err
}

func (p *Plant) importMods() error {
	for _, m := range p.Mods {
		name := m.Name()

		mod, ok := p.mods[name]
		if !ok {
			return errors.New("importMods: mod not added")
		}

		if err := mod.Import(p.Config.V(m)); err != nil {
			return fmt.Errorf("importMods: (%s) %w", name, err)
		}
	}

	return nil
}
