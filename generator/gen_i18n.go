package generator

import (
	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenI18N() error {
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          i18nDir,
		filename:        "i18n.go",
		templateName:    "i18nTemplate",
		builtinTemplate: tpl.I18nTemplate,
	})
}
