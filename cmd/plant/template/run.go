package template

var Run = `
package main

import (
	"log"

	"automatica.team/plant"
	"automatica.team/plant/parts/core"
)

// Import plant deps
import (
	_ "automatica.team/plant/deps/db"
)

func main() {
	// Create a plant
	p, err := plant.New("plant.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to deps
	d, err := p.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Compose a new bot
	b, err := p.Compose()
	if err != nil {
		log.Fatal(err)
	}

	// Add parts to the bot
	p.Add(core.New(b, d))

	// Build the bot
	if err := p.Build(b); err != nil {
		log.Fatal(err)
	}

	// Start the bot
	b.Start()
}
`
