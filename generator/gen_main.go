package generator

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

const ginEngineName = "e"

func GenMain() error {
	filename, err := format.FileNamingFormat(PluginInfo.Style, PluginInfo.Api.Service.Name)
	if err != nil {
		return err
	}

	return genFile(fileGenConfig{
		dir:             PluginInfo.Dir,
		subDir:          "",
		filename:        filename + ".go",
		templateName:    "mainTemplate",
		builtinTemplate: tpl.MainTemplate,
		data: map[string]interface{}{
			"importPkg":     genMainImportPkg(),
			"etcDir":        etcDir,
			"configName":    filename,
			"setup":         genMainSetup(),
			"host":          defaultHost,
			"port":          defaultPort,
			"ginEngineName": ginEngineName,
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
		%s.Setup(%s)
`, routesPacket, ginEngineName)
}
