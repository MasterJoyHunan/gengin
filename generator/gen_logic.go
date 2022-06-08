package generator

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/MasterJoyHunan/go-zero-gin-plugin/prepare"
	"github.com/MasterJoyHunan/go-zero-gin-plugin/tpl"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func GenLogic() error {
	for _, g := range PluginInfo.Api.Service.Groups {
		for _, r := range g.Routes {
			err := genLogicByRoute(g, r)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func genLogicByRoute(group spec.Group, route spec.Route) error {
	logic := getLogicName(route)
	filename, err := format.FileNamingFormat(PluginInfo.Style, logic)
	if err != nil {
		return err
	}

	imports := genLogicImports(route)

	var responseString string
	var returnString string
	var requestString string

	if len(route.RequestTypeName()) > 0 {
		groupNameParse := parseGroupName(typeGroup[route.ResponseTypeName()], typesPacket, typesDir)
		requestString = "req *" + groupNameParse.pkgName + "." + strings.Title(route.RequestTypeName())
	}

	if len(route.ResponseTypeName()) > 0 {
		groupNameParse := parseGroupName(typeGroup[route.ResponseTypeName()], typesPacket, typesDir)
		responseString = "(resp " + groupNameParse.pkgName + "." + strings.Title(route.ResponseTypeName()) + ", err error)"
		returnString = "return"
	} else {
		responseString = "error"
		returnString = "return nil"
	}

	logicGroupNameParse := parseGroupName(group.GetAnnotation(groupProperty), logicPacket, logicDir)
	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          logicGroupNameParse.dirPath,
		filename:        filename + ".go",
		templateName:    "logicTemplate",
		builtinTemplate: tpl.LogicTemplate,
		data: map[string]interface{}{
			"pkgName":      logicGroupNameParse.pkgName,
			"hasImport":    len(imports) > 0,
			"imports":      imports,
			"function":     title.String(strings.TrimSuffix(logic, "Logic")),
			"responseType": responseString,
			"returnString": returnString,
			"request":      requestString,
		},
	})
}

func genLogicImports(route spec.Route) string {
	importSet := collection.NewSet()
	if len(route.RequestTypeName()) > 0 {
		groupNameParse := parseGroupName(typeGroup[route.RequestTypeName()], typesPacket, typesDir)
		importSet.AddStr(fmt.Sprintf("\"%s\"", pathx.JoinPackages(RootPkg, groupNameParse.dirPath)))
	}

	if len(route.ResponseTypeName()) > 0 {
		groupNameParse := parseGroupName(typeGroup[route.ResponseTypeName()], typesPacket, typesDir)
		importSet.AddStr(fmt.Sprintf("\"%s\"", pathx.JoinPackages(RootPkg, groupNameParse.dirPath)))
	}

	importArr := importSet.KeysStr()
	sort.Strings(importArr)
	return strings.Join(importArr, "\n\t")
}
