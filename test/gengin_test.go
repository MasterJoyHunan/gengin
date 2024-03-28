package test

import (
	"testing"

	"github.com/MasterJoyHunan/gengin/generator"
	"github.com/MasterJoyHunan/gengin/prepare"
)

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestGenEtc(t *testing.T) {
	if err := generator.GenEtc(); err != nil {
		t.Failed()
	}
}
func TestGenConfig(t *testing.T) {
	if err := generator.GenConfig(); err != nil {
		t.Failed()
	}
}
func TestGenMain(t *testing.T) {
	if err := generator.GenMain(); err != nil {
		t.Failed()
	}
}
func TestGenMiddleware(t *testing.T) {
	if err := generator.GenMiddleware(); err != nil {
		t.Failed()
	}
}
func TestGenTypes(t *testing.T) {
	if err := generator.GenTypes(); err != nil {
		t.Failed()
	}
}

func TestGenLogic(t *testing.T) {
	if err := generator.GenLogic(); err != nil {
		t.Failed()
	}
}

func TestGenRoutes(t *testing.T) {
	if err := generator.GenRoutes(); err != nil {
		t.Failed()
	}
}

func TestGenHandlers(t *testing.T) {
	if err := generator.GenHandlers(); err != nil {
		t.Failed()
	}
}

func TestGenResponse(t *testing.T) {
	if err := generator.GenResponse(); err != nil {
		t.Failed()
	}
}

func TestGenI18N(t *testing.T) {
	if err := generator.GenI18N(); err != nil {
		t.Failed()
	}
}

func TestGenSvcContext(t *testing.T) {
	if err := generator.GenSvcContext(); err != nil {
		t.Failed()
	}
}

func setup() {
	prepare.OutputDir = "example"
	prepare.ApiFile = "api/someapp.api"
	prepare.Setup()
}
