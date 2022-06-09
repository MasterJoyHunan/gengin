package generator

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
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
	groups := getGinRoutes()
	for _, g := range groups {
		var sb strings.Builder
		for _, r := range g.routes {
			sb.WriteString(fmt.Sprintf("    g.%s(\"%s\", %s)\n", r.method, r.path, r.handler))
		}

		routeFilename, err := format.FileNamingFormat(PluginInfo.Style, routesPacket)
		if err != nil {
			return err
		}

		routeFilename = routeFilename + ".go"
		filename := path.Join(PluginInfo.Dir, g.dirPath, routeFilename)
		os.Remove(filename)

		for i := range g.middlewares {
			g.middlewares[i] = middlewarePacket + "." + g.middlewares[i]
		}

		err = genFile(fileGenConfig{
			dir:             PluginInfo.Dir,
			subDir:          g.dirPath,
			filename:        routeFilename,
			templateName:    "routesTemplate",
			builtinTemplate: tpl.RoutesTemplate,
			data: map[string]interface{}{
				"pkgName":         g.pkgName,
				"prefix":          g.prefix,
				"hasPrefix":       len(g.prefix) > 0,
				"hasMiddleware":   len(g.middlewares) > 0,
				"middleware":      strings.Join(g.middlewares, ", "),
				"function":        util.Title(g.groupName),
				"importPackages":  genGinRouteImports(g),
				"routesAdditions": sb.String(),
			},
		})
		if err != nil {
			return err
		}
	}
	return genSetup(groups)
}

func genGinRouteImports(g groupInfo) string {
	importSet := collection.NewSet()

	//  middleware 的 import
	if len(g.middlewares) > 0 {
		importSet.AddStr(fmt.Sprintf(`"%s"`, pathx.JoinPackages(RootPkg, middlewareDir)))
	}

	//  handle 的 import
	groupNameParse := parseGroupName(g.groupName, handlerDir, handlerPacket)
	importSet.AddStr(fmt.Sprintf(`"%s"`, pathx.JoinPackages(RootPkg, groupNameParse.dirPath)))

	imports := importSet.KeysStr()
	sort.Strings(imports)
	return strings.Join(imports, "\n\t")
}

func getGinRoutes() []groupInfo {
	var groupInfos []groupInfo

	for _, g := range PluginInfo.Api.Service.Groups {
		var groupedRoutes groupInfo
		for _, r := range g.Routes {
			handler := getHandlerName(r)

			routeGroupNameParse := parseGroupName(g.GetAnnotation(groupProperty), routesDir, routesPacket)
			groupedRoutes.groupName = routeGroupNameParse.groupName
			groupedRoutes.dirPath = routeGroupNameParse.dirPath
			groupedRoutes.pkgName = routeGroupNameParse.pkgName

			handlerGroupNameParse := parseGroupName(g.GetAnnotation(groupProperty), handlerDir, handlerPacket)
			groupedRoutes.routes = append(groupedRoutes.routes, route{
				method:  strings.ToUpper(r.Method),
				path:    r.Path,
				handler: handlerGroupNameParse.pkgName + "." + util.Title(handler),
			})
		}

		jwt := g.GetAnnotation("jwt")
		if len(jwt) > 0 {
			groupedRoutes.middlewares = append(groupedRoutes.middlewares, strings.TrimSuffix(jwt, "Middleware")+"Middleware")
		}

		middleware := g.GetAnnotation("middleware")
		if len(middleware) > 0 {
			groupedRoutes.middlewares = append(groupedRoutes.middlewares,
				strings.Split(middleware, ",")...)
		}
		prefix := g.GetAnnotation(spec.RoutePrefixKey)
		prefix = strings.ReplaceAll(prefix, `"`, "")
		prefix = strings.TrimSpace(prefix)
		if len(prefix) > 0 {
			prefix = path.Join("/", prefix)
			groupedRoutes.prefix = prefix
		}
		groupInfos = append(groupInfos, groupedRoutes)
	}

	return groupInfos
}

func genSetup(groups []groupInfo) error {
	filename := "setup.go"
	os.Remove(pathx.JoinPackages(PluginInfo.Dir, routesDir, filename))

	var importArr []string
	var regArr []string
	for _, g := range groups {
		if g.dirPath != routesDir {
			importArr = append(importArr, fmt.Sprintf("\"%s\"", pathx.JoinPackages(RootPkg, g.dirPath)))
			regArr = append(regArr, fmt.Sprintf("%s.Register%sRoute(e)", g.pkgName, util.Title(g.groupName)))
		} else {
			regArr = append(regArr, fmt.Sprintf("Register%sRoute(e)", util.Title(g.groupName)))
		}
	}

	sort.Strings(importArr)
	importStr := strings.Join(importArr, "\n\t")
	regStr := strings.Join(regArr, "\n\t")

	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          routesDir,
		filename:        filename,
		templateName:    "routesTemplate",
		builtinTemplate: tpl.RoutesSetupTemplate,
		data: map[string]interface{}{
			"importPackages": importStr,
			"register":       regStr,
		},
	})
}
