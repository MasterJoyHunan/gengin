package generator

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const labelName = "label"

var requestTypes map[string]any

func GenTypes() error {
	requestTypes = getRequestTypes()
	types, err := buildTypes()
	if err != nil {
		return err
	}

	filename := pathx.JoinPackages(prepare.OutputDir, "types/types.go")
	os.Remove(filename)

	err = GenFile(
		"types.go",
		tpl.TypesTemplate,
		WithSubDir("types"),
		WithData(map[string]any{
			"types": types,
		}),
	)
	if err != nil {
		return err
	}
	return nil
}

func getRequestTypes() map[string]any {
	types := make(map[string]any)
	for _, group := range prepare.ApiSpec.Service.Groups {
		for _, r := range group.Routes {
			types[r.RequestTypeName()] = nil
		}
	}
	return types
}

// buildTypes gen types to string
func buildTypes() (string, error) {
	var builder strings.Builder
	first := true
	for _, tp := range prepare.ApiSpec.Types {
		if first {
			first = false
		} else {
			builder.WriteString("\n\n")
		}
		if err := writeType(&builder, tp); err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}

func writeType(writer io.Writer, tp spec.Type) error {
	structType, ok := tp.(spec.DefineStruct)
	if !ok {
		return fmt.Errorf("unspport struct type: %s", tp.Name())
	}

	fmt.Fprintf(writer, "type %s struct {\n", util.Title(tp.Name()))
	for _, member := range structType.Members {
		if member.IsInline {
			if _, err := fmt.Fprintf(writer, "%s\n", util.Title(member.Type.Name())); err != nil {
				return err
			}
			continue
		}

		tag := OverrideTag(tp, member)

		if err := writeProperty(writer, member.Name, tag, member.GetComment(), member.Type); err != nil {
			return err
		}
	}
	fmt.Fprintf(writer, "}")
	return nil
}

func writeProperty(writer io.Writer, name, tag, comment string, tp spec.Type) error {
	var err error
	if len(comment) > 0 {
		comment = strings.TrimPrefix(comment, "//")
		comment = "//" + comment
		_, err = fmt.Fprintf(writer, "%s %s %s %s\n", util.Title(name), tp.Name(), tag, comment)
	} else {
		_, err = fmt.Fprintf(writer, "%s %s %s\n", util.Title(name), tp.Name(), tag)
	}
	return err
}

func OverrideTag(tp spec.Type, member spec.Member) string {
	// 将 path 替换为 uri
	tag := member.Tag
	before, _, found := strings.Cut(tag, ":")
	if found && strings.HasSuffix(before, "path") {
		tag = strings.Replace(tag, "path", "uri", 1)
	}

	// 将注释加入到 label, 用于 validator 验证时中文返回 see http://github.com/go-playground/validator/v10
	// 希望只对 request type 进行处理
	_, ok := requestTypes[tp.Name()]
	if !ok {
		return tag
	}

	label := ""
	if member.Comment != "" {
		label = strings.ReplaceAll(member.Comment, "/", "")
		label = strings.Trim(label, " ")
	}
	if label != "" {
		label = fmt.Sprintf("%s:\"%s\"", labelName, label)
		tag = fmt.Sprintf("%s %s`", tag[:len(tag)-1], label)
	}
	return tag
}
