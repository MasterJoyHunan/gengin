package main

import (
	"flag"

	"github.com/MasterJoyHunan/gengin/test/example/routes"

	"github.com/gin-gonic/gin"
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	// configFile := fmt.Sprintf("etc/someapp-%s.yaml", release)

	e := gin.Default()

	routes.Setup(e)

	e.Run("127.0.0.1:8888")
}
