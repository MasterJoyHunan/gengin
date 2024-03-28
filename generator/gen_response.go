package generator

import (
	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenResponse() error {
	return GenFile(
		"response.go",
		tpl.ResponseTemplate,
		WithSubDir("internal/response"),
		WithData(map[string]string{
			"rootPkg": prepare.RootPkg,
		}),
	)
}
