package prepare

import (
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util/ctx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	ApiSpec   *spec.ApiSpec
	RootPkg   string
	OutputDir string
	ApiFile   string
)

func Setup() {
	var err error
	ApiSpec, err = parser.Parse(ApiFile)
	if err != nil {
		panic(err)
	}

	if err = ApiSpec.Validate(); err != nil {
		panic(err)
	}

	RootPkg, err = GetParentPackage(OutputDir)
	if err != nil {
		panic(err)
	}
}

func GetParentPackage(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	projectCtx, err := ctx.Prepare(abs)
	if err != nil {
		return "", err
	}

	wd := projectCtx.WorkDir
	d := projectCtx.Dir
	same, err := pathx.SameFile(wd, d)
	if err != nil {
		return "", err
	}

	trim := strings.TrimPrefix(projectCtx.WorkDir, projectCtx.Dir)
	if same {
		trim = strings.TrimPrefix(strings.ToLower(projectCtx.WorkDir), strings.ToLower(projectCtx.Dir))
	}

	return filepath.ToSlash(filepath.Join(projectCtx.Path, trim)), nil
}
