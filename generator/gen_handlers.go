package generator

import (
	"path"
	"strings"

	"github.com/MasterJoyHunan/gengin/pkg"
	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenHandlers() error {
	for _, group := range prepare.ApiSpec.Service.Groups {
		for _, r := range group.Routes {
			if err := genHandler(group, r); err != nil {
				return err
			}
		}
	}
	return nil
}

func genHandler(group spec.Group, route spec.Route) error {
	handlerName, err := format.FileNamingFormat(fileNameStyle, route.Handler)
	if err != nil {
		return err
	}

	handlerFileName := strings.TrimSuffix(strings.TrimSuffix(handlerName, "handle"), "_") + "_handle.go"

	subDir := group.GetAnnotation(groupProperty)
	subDir, err = format.FileNamingFormat(dirStyle, subDir)
	if err != nil {
		return err
	}

	handlePkg := path.Join("handler", subDir)
	logicPkg := path.Join("logic", subDir)

	handleBase := path.Base(handlePkg)
	logicBase := path.Base(logicPkg)

	// 解析请求
	p := new(pkg.ParseRequestBody)
	parseRequest := p.BuildParseRequestStr(route.RequestType)

	return GenFile(
		handlerFileName,
		tpl.HandlerTemplate,
		WithSubDir(handlePkg),
		WithData(map[string]any{
			"rootPkg":      prepare.RootPkg,
			"pkgName":      handleBase,
			"logicPkg":     logicPkg,
			"logicBase":    logicBase,
			"comment":      parseComment(route),
			"handlerName":  cases.Title(language.English, cases.NoLower).String(route.Handler),
			"requestType":  cases.Title(language.English, cases.NoLower).String(route.RequestTypeName()),
			"hasResp":      len(route.ResponseTypeName()) > 0,
			"hasReq":       len(route.RequestTypeName()) > 0,
			"parseRequest": parseRequest,
		}),
	)
}
