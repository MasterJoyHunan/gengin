package generator

import (
	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenMain() error {
	filename, err := format.FileNamingFormat(fileNameStyle, prepare.ApiSpec.Service.Name)
	if err != nil {
		return err
	}

	return GenFile(
		filename+".go",
		tpl.MainTemplate,
		WithData(map[string]interface{}{
			"rootPkg":    prepare.RootPkg,
			"configName": filename,
			"host":       defaultHost,
			"port":       defaultPort,
		}),
	)
}
