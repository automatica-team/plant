package main

import (
	"automatica.team/plant"
	"automatica.team/plant/deps/db"
	"automatica.team/plant/mods/core"
)

func main() {
	p, err := plant.New("plant.yml")
	if err != nil {
		panic(err)
	}

	// Dependencies
	p.Inject(&db.DB{})

	// Modules
	p.Add(&core.Core{})

	// Build a configured bot
	b, err := p.Build()
	if err != nil {
		panic(err)
	}

	// Start the built bot
	b.Start()
}
