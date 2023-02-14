package generator

import (
	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenSvcContext() error {
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          svcDir,
		filename:        svcPacket + "_context.go",
		templateName:    "svcTemplate",
		builtinTemplate: tpl.SvcContextTemplate,
		data: map[string]interface{}{
			"pkgName": svcPacket,
		},
	})
}
