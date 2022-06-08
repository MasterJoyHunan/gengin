package generator

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/MasterJoyHunan/go-zero-gin-plugin/prepare"
	"github.com/MasterJoyHunan/go-zero-gin-plugin/tpl"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

func GenMain() error {
	filename, err := format.FileNamingFormat(PluginInfo.Style, PluginInfo.Api.Service.Name)
	if err != nil {
		return err
	}

	if strings.HasSuffix(filename, "-api") {
		filename = strings.ReplaceAll(filename, "-api", "")
	}

	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		builtinTemplate: tpl.MainTemplate,
		data: map[string]interface{}{
			"importPkg":  genMainImportPkg(),
			"etcDir":     etcDir,
			"configName": filename,
			"setup":      genMainSetup(),
			"host":       defaultHost,
			"port":       defaultPort,
		},
	})
}

func genMainImportPkg() string {
	set := collection.NewSet()

	// route
	set.AddStr(fmt.Sprintf("\"%s/%s\"", RootPkg, routesDir))

	// config
	//set.AddStr(fmt.Sprintf())

	importArr := set.KeysStr()
	sort.Strings(importArr)
	return strings.Join(importArr, "\n\t")
}

func genMainSetup() string {
	return fmt.Sprintf(`
		%s.Setup()
`, routesPacket)
}
