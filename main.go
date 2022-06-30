package main

import (
	"github.com/MasterJoyHunan/gengin/generator"
	"github.com/MasterJoyHunan/gengin/prepare"
)

func main() {
	prepare.Setup()
	Must(generator.GenEtc())
	Must(generator.GenConfig())
	Must(generator.GenMain())
	Must(generator.GenMiddleware())
	Must(generator.GenTypes())
	Must(generator.GenLogic())
	Must(generator.GenRoutes())
	Must(generator.GenHandlers())
	Must(generator.GenResponse())
	Must(generator.GenI18N())
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
