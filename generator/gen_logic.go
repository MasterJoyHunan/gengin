package generator

import (
	"path"
	"strings"

	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenLogic() error {
	for _, g := range prepare.ApiSpec.Service.Groups {
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
	logicName, err := format.FileNamingFormat(fileNameStyle, route.Handler)
	if err != nil {
		return err
	}

	logicFileName := strings.TrimSuffix(strings.TrimSuffix(logicName, "logic"), "_") + "_logic.go"

	subDir := group.GetAnnotation(groupProperty)
	subDir, err = format.FileNamingFormat(dirStyle, subDir)
	if err != nil {
		return err
	}

	logicPkg := path.Join("logic", subDir)
	logicBase := path.Base(logicPkg)

	respIsPrimitiveType, respTypeName := parseResponseType(route.ResponseType)

	return GenFile(
		logicFileName,
		tpl.LogicTemplate,
		WithSubDir(logicPkg),
		WithData(map[string]any{
			"rootPkg":           prepare.RootPkg,
			"pkgName":           logicBase,
			"comment":           parseComment(route),
			"logicName":         cases.Title(language.English, cases.NoLower).String(route.Handler),
			"requestType":       cases.Title(language.English, cases.NoLower).String(route.RequestTypeName()),
			"responseType":      respTypeName,
			"needImportTypePkg": len(route.RequestTypeName()) > 0 || (!respIsPrimitiveType && len(route.ResponseTypeName()) > 0),
			"hasReq":            len(route.RequestTypeName()) > 0,
			"hasResp":           len(route.ResponseTypeName()) > 0,
		}),
	)
}
