package tmpl

import (
	"text/template"
)

var Run = func() *template.Template {
	t := template.New("Run").Funcs(funcs)
	t = template.Must(t.New("main.go").Parse(runMainGo))
	t = template.Must(t.New("bot.yml").Parse(runBotYml))
	return t
}()

const runMainGo = "//" + genHeader + `
package main

import (
	"log"
)

import (
	"automatica.team/plant"

	{{- range .OfPrefix .Deps "plant/" }}
	"automatica.team/{{ import .Name "deps" }}"
	{{- end }}

	{{- range .OfPrefix .Mods "plant/" }}
	"automatica.team/{{ import .Name "mods" }}"
	{{- end }}
)

import (
	{{- range .OfPrefix .Deps "x/" }}
	"{{ $.ModName }}/deps/{{ basename .Name }}"
	{{- end }}

	{{- range .OfPrefix .Mods "x/" }}
	"{{ $.ModName }}/mods/{{ basename .Name }}"
	{{- end }}
)

func main() {
	// Create a plant
	p, err := plant.New("plant.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Dependencies
	{{- range .Deps }}
	p.Inject({{ basename .Name }}.New())
	{{- end }}

	// Modules
	{{- range .Mods }}
	p.Add({{ basename .Name }}.New())
	{{- end }}

	// Build a configured bot
	b, err := p.Build()
	if err != nil {
		log.Fatal(err)
	}

	// Start the built bot
	b.Start()
}
`

const runBotYml = "#" + genHeader + `
settings:
  token: {{ .Bot.Token.String }}
  parse_mode: html
`