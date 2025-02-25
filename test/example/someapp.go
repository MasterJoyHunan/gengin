package main

import (
	"flag"
	"fmt"
	"net/http"

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

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", "127.0.0.1", 8888),
		Handler: e,
	}

	server.ListenAndServe()
}
