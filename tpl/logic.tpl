package {{.pkgName}}

import (
	"{{.rootPkg}}/svc"
	{{if or .hasResp .hasResp}}"{{.rootPkg}}/types"{{end}}
)

// {{.logicName}} {{.comment}}
func {{.logicName}}(ctx *svc.ServiceContext{{if .hasReq}}, req *types.{{.requestType}}{{end}}) {{if .hasResp}}(resp types.{{.responseType}}, err error){{else}}error{{end}} {
	// todo: add your logic here and delete this line

	{{if .hasResp}}return{{else}}return nil{{end}}
}
