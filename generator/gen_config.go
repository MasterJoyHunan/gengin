package generator

import (
	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
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
