package generator

import (
	"fmt"
	"strings"

	. "github.com/MasterJoyHunan/go-zero-gin-plugin/prepare"
	"github.com/MasterJoyHunan/go-zero-gin-plugin/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 8888
	devModel    = "local|dev|prod"
)

func GenEtc() error {
	filename, err := format.FileNamingFormat(PluginInfo.Style, PluginInfo.Api.Service.Name)
	if err != nil {
		return err
	}

	mode := strings.Split(devModel, "|")

	for _, m := range mode {
		err = genFile(fileGenConfig{
			dir:             PluginInfo.Dir,
			subDir:          etcDir,
			filename:        fmt.Sprintf("%s-%s.yaml", filename, m),
			templateName:    "etcTemplate",
			builtinTemplate: tpl.EtcTemplate,
			data: map[string]interface{}{
				"serviceName": PluginInfo.Api.Service.Name,
				"host":        defaultHost,
				"port":        defaultPort,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
