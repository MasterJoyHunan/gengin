package generator

import (
	"bytes"
	goformat "go/format"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/api/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

type (
	fileGenConfig struct {
		dir             string
		subDir          string
		filename        string
		templateName    string
		builtinTemplate string
		data            interface{}
	}

	groupBase struct {
		groupName string
		dirPath   string
		pkgName   string
	}
)

func genFile(c fileGenConfig) error {
	fp, created, err := util.MaybeCreateFile(c.dir, c.subDir, c.filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	// 暂时不支持自定义模板
	text := c.builtinTemplate

	t := template.Must(template.New(c.templateName).Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.data)
	if err != nil {
		return err
	}

	code := formatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func formatCode(code string) string {
	ret, err := goformat.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
}

// parseGroupName 解析 Group 所属的 subDir 和 pkgName
func parseGroupName(groupName, defaultDir, defaultPkgName string) (i groupBase) {
	if groupName == "" {
		i.groupName = ""
		i.dirPath = defaultDir
		i.pkgName = defaultPkgName
		return
	}
	fmtName, err := format.FileNamingFormat(dirFmt, groupName)
	logx.Must(err)

	i.groupName = groupName
	i.dirPath = pathx.JoinPackages(defaultDir, fmtName)
	i.pkgName = fmtName[strings.LastIndex(fmtName, "/")+1:]
	return
}

func getHandlerBaseName(route spec.Route) string {
	handler := route.Handler
	handler = strings.TrimSpace(handler)
	handler = strings.TrimSuffix(handler, "handler")
	handler = strings.TrimSuffix(handler, "Handler")
	return handler
}

func getLogicName(route spec.Route) string {
	return getHandlerBaseName(route) + "Logic"
}

func getHandlerName(route spec.Route) string {
	return getHandlerBaseName(route) + "Handle"
}

func getTypesImportAlias(pkg groupBase) string {
	if pkg.dirPath == typesPacket {
		return ""
	}
	return pkg.pkgName + typePkgAlias + " "
}

func getTypesUseAlias(pkg groupBase) string {
	if pkg.dirPath == typesPacket {
		return typesPacket + "."
	}
	return pkg.pkgName + typePkgAlias + "."
}

func parseComment(r spec.Route) string {
	if r.AtDoc.Text != "" {
		return strings.Trim(r.AtDoc.Text, "\"")
	}
	if len(r.HandlerDoc) != 0 {
		str := ""
		for _, d := range r.HandlerDoc {
			str += strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(d, "/", ""), "*", ""), "\n", ""), "\t", "")
		}
		return str
	}
	return ""
}
