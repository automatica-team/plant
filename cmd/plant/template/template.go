package template

import (
	"strings"
	"text/template"
)

var funcs = template.FuncMap{
	"importmods": func(s string) string {
		return strings.ReplaceAll(s, "/", "/mods/")
	},
	"importdeps": func(s string) string {
		return strings.ReplaceAll(s, "/", "/deps/")
	},
	"modname": func(s string) string {
		return s[strings.LastIndex(s, "/")+1:]
	},
}
