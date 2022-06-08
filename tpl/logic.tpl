package {{.pkgName}}

{{if .hasImport}}import (
	{{.imports}}
){{end}}


func {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
