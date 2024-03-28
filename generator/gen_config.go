package generator

import (
	"github.com/MasterJoyHunan/gengin/tpl"
)

func GenConfig() error {
	return GenFile(
		"config.go",
		tpl.ConfigTemplate,
		WithSubDir("config"),
	)
}
