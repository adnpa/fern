package generate

const (
	OPTION_TMPL = `
// Code generated by astgen; DO NOT EDIT.

package config

{{range .}}
// With{{.Name}} with {{.Name}}.
func With{{.Name}}({{.ParamName}} {{.Type}}) Option {
	return func(o *options) {
		o.{{.Name}} = {{.ParamName}}
	}
}
{{end}}
`


)
