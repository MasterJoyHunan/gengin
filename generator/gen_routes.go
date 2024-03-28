package generator

import (
	"os"
	"path"
	"strings"

	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

type (
	groupInfo struct {
		routes      []route
		middlewares []string
		prefix      string
		groupBase
	}
	route struct {
		method  string
		path    string
		handler string
	}
)

func GenRoutes() error {
	for _, group := range prepare.ApiSpec.Service.Groups {
		subDir := group.GetAnnotation(groupProperty)
		subDir, err := format.FileNamingFormat(dirStyle, subDir)
		if err != nil {
			return err
		}

		routesPkg := path.Join("routes", subDir)
		routesBase := path.Base(routesPkg)

		os.Remove(path.Join(prepare.OutputDir, routesPkg, "route.go"))

		// handle
		handlePkg := path.Join("handler", subDir)
		handleBase := path.Base(handlePkg)

		// prefix
		prefix := group.GetAnnotation(spec.RoutePrefixKey)

		// middlewares
		var middlewares []string
		if len(group.GetAnnotation("jwt")) > 0 {
			middlewares = append(middlewares, group.GetAnnotation("jwt"))
		}

		if len(group.GetAnnotation("middleware")) > 0 {
			middlewares = append(middlewares, strings.Split(group.GetAnnotation("middleware"), ",")...)
		}

		middlewares = lo.Map(middlewares, func(item string, index int) string {
			res := strings.TrimSuffix(item, "Middleware") + "Middleware"
			return cases.Title(language.English, cases.NoLower).String(res)
		})

		// route
		var routes []map[string]string
		for _, r := range group.Routes {
			routes = append(routes, map[string]string{
				"method": strings.ToUpper(r.Method),
				"path":   r.Path,
				"handle": cases.Title(language.English, cases.NoLower).String(r.Handler),
			})
		}

		err = GenFile(
			"route.go",
			tpl.RoutesTemplate,
			WithSubDir(routesPkg),
			WithData(map[string]any{
				"rootPkg":       prepare.RootPkg,
				"pkgName":       routesBase,
				"handlePkg":     handlePkg,
				"handleBase":    handleBase,
				"prefix":        prefix,
				"hasPrefix":     len(prefix) > 0,
				"hasMiddleware": len(middlewares) > 0,
				"middleware":    middlewares,
				"funcName":      cases.Title(language.English, cases.NoLower).String(group.GetAnnotation(groupProperty)),
				"routes":        routes,
			}),
		)

		if err != nil {
			return err
		}
	}
	return genSetup()
}

func genSetup() error {
	os.Remove(path.Join(prepare.OutputDir, "routes/setup.go"))

	var routes []map[string]string
	for _, group := range prepare.ApiSpec.Service.Groups {
		subDir := group.GetAnnotation(groupProperty)
		subDir, err := format.FileNamingFormat(dirStyle, subDir)
		if err != nil {
			return err
		}

		routesPkg := path.Join("routes", subDir)
		routesBase := path.Base(routesPkg)

		routes = append(routes, map[string]string{
			"pkg":  routesPkg,
			"base": routesBase,
			"name": cases.Title(language.English, cases.NoLower).String(group.GetAnnotation(groupProperty)),
		})

	}

	return GenFile(
		"setup.go",
		tpl.RoutesSetupTemplate,
		WithSubDir("routes"),
		WithData(map[string]any{
			"rootPkg": prepare.RootPkg,
			"routes":  routes,
		}),
	)
}
