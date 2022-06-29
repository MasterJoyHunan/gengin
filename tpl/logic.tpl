package {{.pkgName}}

{{if .hasImport}}import (
	{{.imports}}
){{end}}

// {{.function}} {{.comment}}
func {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line

	{{.returnString}}
}
