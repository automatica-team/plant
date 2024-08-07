package private

import (
	"automatica.team/plant"
)

func (mod *Private) Name() string {
	return "x/private"
}

type Private struct {
	plant.Handler
}

func New() *Private {
	return &Private{}
}

func (mod *Private) Import(m plant.M) error {
	return nil
}
