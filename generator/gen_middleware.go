package generator

import (
	"strings"

	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenMiddleware() error {
	middlewares := getMiddleware()
	for _, item := range middlewares {
		middlewareFilename := strings.TrimSuffix(strings.ToLower(item), "middleware") + "_middleware"
		filename, err := format.FileNamingFormat(PluginInfo.Style, middlewareFilename)
		if err != nil {
			return err
		}

		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		err = genFile(fileGenConfig{
			dir:             PluginInfo.Dir,
			subDir:          middlewareDir,
			filename:        filename + ".go",
			templateName:    "contextTemplate",
			builtinTemplate: tpl.MiddlewareTemplate,
			data: map[string]string{
				"name": title.String(name),
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getMiddleware() []string {
	result := collection.NewSet()
	for _, g := range PluginInfo.Api.Service.Groups {
		middleware := g.GetAnnotation("middleware")
		if len(middleware) > 0 {
			for _, item := range strings.Split(middleware, ",") {
				result.Add(strings.TrimSpace(item))
			}
		}
		jwtMiddleware := g.GetAnnotation("jwt")
		if len(jwtMiddleware) > 0 {
			result.Add(strings.TrimSpace(jwtMiddleware))
		}
	}

	return result.KeysStr()
}
