package template

import (
	"strings"
	"text/template"
)

var funcs = template.FuncMap{
	"import": func(s, what string) string {
		return strings.ReplaceAll(s, "/", "/"+what+"/")
	},
	"basename": func(s string) string {
		return s[strings.LastIndex(s, "/")+1:]
	},
}
