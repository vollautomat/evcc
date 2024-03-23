package {{.Package}}

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"{{.API}}"
)

{{define "case"}}
	{{- $combo := .Combo}}
	{{- $prefix := .Prefix}}
	{{- $idx := 0}}

	{{- range $typ, $def := .Types}}
		{{- range $method := $def.Methods }}
			{{- if gt $idx 0}} &&{{else}}{{$idx = 1}}{{end}} {{$method.VarName}} {{if contains $combo $typ}}!={{else}}=={{end}} nil
		{{- end}}
	{{- end}}:
		return &struct {
			{{.BaseType}}
{{- range $typ, $def := .Types}}
	{{- if contains $combo $typ}}
			{{$typ}}
	{{- end}}
{{- end}}
		}{
			{{.ShortBase}}: base,
{{- range $typ, $def := .Types}}
	{{- if contains $combo $typ}}
			{{$def.ShortType}}: &{{$prefix}}{{$def.ShortType}}Impl{
				{{- range $method := $def.Methods }}
					{{$method.VarName}}: {{$method.VarName}},
				{{- end}}
			},
	{{- end}}
{{- end}}
		}
{{- end -}}

func {{.Function}}(base {{.BaseType}}{{range $type := ordered}}{{range $m := $type.Methods}}, {{$m.VarName}} {{$m.Signature}}{{end}}{{end}}) {{.ReturnType}} {
{{- $basetype := .BaseType}}
{{- $shortbase := .ShortBase}}
{{- $prefix := .Function}}
{{- $types := .Types}}
{{- $idx := 0}}
	switch {
	case {{- range $typ, $def := .Types}}
		{{- range $method := $def.Methods}}
			{{- if gt $idx 0}} &&{{else}}{{$idx = 1}}{{end}} {{$method.VarName}} == nil
		{{- end}}
	{{- end}}:
		return base
{{range $combo := .Combinations}}
	case {{- template "case" dict "BaseType" $basetype "Prefix" $prefix "ShortBase" $shortbase "Types" $types "Combo" $combo}}
{{end}}	}

	return nil
}

{{range $t := .Types -}}
type {{$prefix}}{{.ShortType}}Impl struct {
	{{range .Methods -}}
		{{.VarName}} {{.Signature}}
	{{end}}
}
	{{range $m := .Methods -}}
	func (impl *{{$prefix}}{{$t.ShortType}}Impl) {{.Function}}(
		{{- range $idx, $param := .Params -}}
			{{- if gt $idx 0}}, {{end -}}
			p{{$idx}} {{ $param -}} 
		{{end}}){{ .ReturnTypes }} {
		return impl.{{$m.VarName}}(
		{{- range $idx, $param := .Params -}}
			{{- if gt $idx 0}}, {{end}}p{{- $idx -}}{{end}})
		}
	{{end}}
{{end}}
