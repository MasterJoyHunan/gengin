package generator

import (
	"bytes"
	goformat "go/format"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type fileGenConfig struct {
	dir          string
	subDir       string
	filename     string
	templateName string
	templateText string
	data         any
}

func GenFile(fileName, templateText string, opt ...Option) error {
	templateName, _, _ := strings.Cut(fileName, ".")

	cfg := &fileGenConfig{
		filename:     fileName,
		templateName: templateName,
		templateText: templateText,
	}
	for _, fn := range opt {
		fn(cfg)
	}

	if len(cfg.dir) == 0 {
		cfg.dir = prepare.OutputDir
	}

	filePath := path.Join(cfg.dir, cfg.subDir, cfg.filename)
	_, err := os.Stat(filePath)
	if err == nil {
		// 文件已存在
		return nil
	}

	err = os.MkdirAll(path.Join(cfg.dir, cfg.subDir), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.New(cfg.templateName).Parse(cfg.templateText))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, cfg.data)
	if err != nil {
		return err
	}

	code := formatCode(buffer.String())
	_, err = file.WriteString(code)

	return err
}

type Option func(*fileGenConfig)

// WithDir 设置目录
func WithDir(dir string) Option {
	return func(config *fileGenConfig) {
		config.dir = dir
	}
}

// WithSubDir 设置二级目录
func WithSubDir(dir string) Option {
	return func(config *fileGenConfig) {
		config.subDir = dir
	}
}

// WithData 设置数据
func WithData(data any) Option {
	return func(config *fileGenConfig) {
		config.data = data
	}
}

func formatCode(code string) string {
	ret, err := goformat.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
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

func parseResponseType(t spec.Type) (isPrimitiveType bool, typeName string) {
	if t == nil {
		return true, ""
	}
	switch v := t.(type) {
	case spec.DefineStruct:
		return false, "types." + cases.Title(language.English, cases.NoLower).String(t.Name())
	case spec.PrimitiveType: // 内置
		return true, t.Name()
	case spec.MapType:
		// 不允许
	case spec.ArrayType:
		isPrimitiveType, typeName = parseResponseType(v.Value)
		return isPrimitiveType, "[]" + typeName
	case spec.InterfaceType:
		// 不允许
	case spec.PointerType:
		// 不允许
	}
	return true, ""
}
