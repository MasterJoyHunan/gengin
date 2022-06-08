package main

import (
	"github.com/MasterJoyHunan/go-zero-gin-plugin/generator"
	"github.com/MasterJoyHunan/go-zero-gin-plugin/prepare"
)

func main() {
	prepare.Setup()
	Must(generator.GenEtc())
	Must(generator.GenConfig())
	Must(generator.GenMain())
	Must(generator.GenMiddleware())
	Must(generator.GenTypes())
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
