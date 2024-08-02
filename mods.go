package plant

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"unsafe"
)

type Mod interface {
	Name() string
	Expose() []string
	Import(M) error
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
				dep, ok := p.deps[strings.TrimPrefix(tag, "dep:")]
				if !ok {
					continue
				}

				uf.Set(reflect.ValueOf(dep))
			}
		}
	}

	return err
}

func (p *Plant) importMods() error {
	//if len(p.Mods) == 0 {
	//	return errors.New("importMods: empty mods")
	//}

	for _, m := range p.Mods {
		name := m.Name()

		mod, ok := p.mods[name]
		if !ok {
			return errors.New("importMods: mod not added")
		}

		if err := mod.Import(m); err != nil {
			return fmt.Errorf("importMods: (%s) %w", name, err)
		}
	}

	return nil
}

func (p *Plant) filterMods(end string) []Mod {
	var mods []Mod
	for _, m := range p.mods {
		if slices.Contains(m.Expose(), end) {
			mods = append(mods, m)
		}
	}
	return mods
}