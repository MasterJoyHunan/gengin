package generator

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	goformat "go/format"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/api/util"
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

	GroupInfo struct {
		GroupName string
		DirPath   string
		PkgName   string
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
func parseGroupName(groupName, defaultDir, defaultPkgName string) (i GroupInfo) {
	if groupName == "" {
		i.GroupName = ""
		i.DirPath = defaultDir
		i.PkgName = defaultPkgName
		return
	}
	fmtName, err := format.FileNamingFormat(dirFmt, groupName)
	logx.Must(err)

	i.GroupName = groupName
	i.DirPath = pathx.JoinPackages(defaultDir, fmtName)
	i.PkgName = fmtName[strings.LastIndex(fmtName, "/")+1:]
	return
}
