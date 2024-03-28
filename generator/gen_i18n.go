package generator

import (
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenI18N() error {
	return GenFile(
		"i18n.go",
		tpl.I18nTemplate,
		WithSubDir("internal/translator"),
	)
}
