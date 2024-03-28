package main

import (
	"flag"
	"fmt"
	"net/http"

	"{{.rootPkg}}/routes"

	"github.com/gin-gonic/gin"
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	// configFile := fmt.Sprintf("etc/{{.configName}}-%s.yaml", release)

    e := gin.Default()

	routes.Setup(e)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", "{{.host}}", {{.port}}),
		Handler: e,
	}

	server.ListenAndServe()
}
