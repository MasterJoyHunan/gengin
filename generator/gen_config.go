package generator

import (
	. "github.com/MasterJoyHunan/go-zero-gin-plugin/prepare"
	"github.com/MasterJoyHunan/go-zero-gin-plugin/tpl"
)

func GenConfig() error {
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          configDir,
		filename:        configPacket + ".go",
		templateName:    "configTemplate",
		builtinTemplate: tpl.ConfigTemplate,
		data: map[string]interface{}{
			"pkgName": configPacket,
		},
	})
}
