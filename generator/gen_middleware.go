package generator

import (
	"strings"

	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
	"github.com/samber/lo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenMiddleware() error {
	middlewares := getMiddlewares()
	for _, item := range middlewares {
		middlewareName, err := format.FileNamingFormat(fileNameStyle, item)
		if err != nil {
			return err
		}
		filename := strings.TrimSuffix(strings.TrimSuffix(strings.ToLower(middlewareName), "middleware"), "_") + "_middleware"

		name := strings.TrimSuffix(item, "Middleware") + "Middleware"

		err = GenFile(
			filename+".go",
			tpl.MiddlewareTemplate,
			WithSubDir("middleware"),
			WithData(map[string]string{
				"name": cases.Title(language.English, cases.NoLower).String(name),
			}),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func getMiddlewares() []string {
	middlewares := make(map[string]any)
	for _, g := range prepare.ApiSpec.Service.Groups {
		middleware := g.GetAnnotation("middleware")
		if len(middleware) > 0 {
			for _, item := range strings.Split(middleware, ",") {
				middlewares[strings.TrimSpace(item)] = nil
			}
		}
		jwtMiddleware := g.GetAnnotation("jwt")
		if len(jwtMiddleware) > 0 {
			middlewares[strings.TrimSpace(jwtMiddleware)] = nil
		}
	}

	return lo.Keys(middlewares)
}
