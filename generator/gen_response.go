package generator

import (
	"fmt"
	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func GenResponse() error {
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          responseDir,
		filename:        responsePacket + ".go",
		templateName:    "responseTemplate",
		builtinTemplate: tpl.ResponseTemplate,
		data: map[string]interface{}{
			"importPkg": fmt.Sprintf("\"%s\"", pathx.JoinPackages(RootPkg, i18nDir)),
		},
	})
}
