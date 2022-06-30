package generator

import (
	"strings"

	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenI18N() error {
	filename, err := format.FileNamingFormat(PluginInfo.Style, PluginInfo.Api.Service.Name)
	if err != nil {
		return err
	}

	if strings.HasSuffix(filename, "-api") {
		filename = strings.ReplaceAll(filename, "-api", "")
	}

	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		builtinTemplate: tpl.MainTemplate,
		data: map[string]interface{}{
			"importPkg":     genMainImportPkg(),
			"etcDir":        etcDir,
			"configName":    filename,
			"setup":         genMainSetup(),
			"host":          defaultHost,
			"port":          defaultPort,
			"ginEngineName": ginEngineName,
		},
	})
}
