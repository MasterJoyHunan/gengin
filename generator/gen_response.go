package generator

import (
	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenResponse() error {
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          responseDir,
		filename:        responsePacket + ".go",
		templateName:    "responseTemplate",
		builtinTemplate: tpl.ResponseTemplate,
	})
}
