package generator

import (
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenSvcContext() error {
	return GenFile(
		"svc_context.go",
		tpl.SvcContextTemplate,
		WithSubDir("svc"),
	)
}
