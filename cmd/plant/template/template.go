package template

import (
	"strings"
	"text/template"
)

var funcs = template.FuncMap{
	"importpart": func(s string) string {
		return strings.ReplaceAll(s, "/", "/parts/")
	},
	"importdeps": func(s string) string {
		return strings.ReplaceAll(s, "/", "/deps/")
	},
	"partname": func(s string) string {
		return s[strings.LastIndex(s, "/")+1:]
	},
}
