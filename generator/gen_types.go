package generator

import (
	"fmt"
	"io"
	"os"
	"strings"

	. "github.com/MasterJoyHunan/gengin/prepare"
	"github.com/MasterJoyHunan/gengin/tpl"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var groupTypes []TypeBelongGroup
var typeGroup = make(map[string]string)
var requestTypes = make(map[string]int)

const labelName = "label"

type TypeBelongGroup struct {
	GroupName string
	TypeStr   string
	TypeMap   []spec.Type
}

func GenTypes() error {
	types, err := BuildGroupTypes()
	if err != nil {
		return err
	}

	typeFilename := typesPacket + ".go"

	for _, t := range types {
		typeGroupInfo := parseGroupName(t.GroupName, typesDir, typesPacket)
		filename := pathx.JoinPackages(PluginInfo.Dir, typeGroupInfo.dirPath, typeFilename)
		os.Remove(filename)

		err = genFile(fileGenConfig{
			dir:             PluginInfo.Dir,
			subDir:          typeGroupInfo.dirPath,
			filename:        typeFilename,
			templateName:    "typesTemplate",
			builtinTemplate: tpl.TypesTemplate,
			data: map[string]interface{}{
				"pkgName": typeGroupInfo.pkgName,
				"types":   t.TypeStr,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// buildTypes gen types to string
func buildTypes(types []spec.Type) (string, error) {
	var builder strings.Builder
	first := true
	for _, tp := range types {
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

// BuildGroupTypes gen types to string
func BuildGroupTypes() ([]TypeBelongGroup, error) {
	// 用于保存 type 被哪几个 groupInfo 用到
	container := make(map[string]map[string]int, 0)
	for _, group := range PluginInfo.Api.Service.Groups {
		for _, route := range group.Routes {
			joinContainer(container, route.RequestType, group.GetAnnotation(groupProperty), true)
			joinContainer(container, route.ResponseType, group.GetAnnotation(groupProperty), false)
		}
	}

	// 用于保存 groupInfo 下包含几个 type， 如果一个 type 被多个 groupInfo 用到了，则放入公共的 type 文件中
	// 如果一个 type 没有设定 groupInfo ，则也会放到 公共的 type 文件中
	groupType := make(map[string]map[string]int)
	for typeName, groups := range container {
		_, ok := groups[""]
		groupName := ""
		if !ok && len(groups) == 1 {
			for group := range groups {
				groupName = group
			}
		}
		_, ok = groupType[groupName]
		if !ok {
			groupType[groupName] = make(map[string]int)
		}
		groupType[groupName][typeName] = 1
	}

	// 包装返回
	for group, typeNames := range groupType {
		var temp []spec.Type
		for _, t := range PluginInfo.Api.Types {
			_, ok := typeNames[t.Name()]
			if ok {
				temp = append(temp, t)
			}
		}
		typeStr, err := buildTypes(temp)
		if err != nil {
			return nil, err
		}
		groupTypes = append(groupTypes, TypeBelongGroup{
			GroupName: group,
			TypeStr:   typeStr,
			TypeMap:   temp,
		})
	}

	// 每个 type 对应的 groupInfo
	for _, g := range groupTypes {
		for _, s := range g.TypeMap {
			typeGroup[s.Name()] = g.GroupName
		}
	}

	return groupTypes, nil
}

func joinContainer(container map[string]map[string]int, defType spec.Type, group string, isRequestType bool) {
	defineStruct, ok := defType.(spec.DefineStruct)
	if !ok {
		return
	}
	for _, t := range PluginInfo.Api.Types {
		if t.Name() == defType.Name() {
			defineStruct = t.(spec.DefineStruct)
		}
	}

	typeName := defineStruct.Name()

	if isRequestType {
		requestTypes[typeName] = 1
	}

	members := defineStruct.Members
	for _, m := range members {
		switch v := m.Type.(type) {
		case spec.MapType:
			joinContainer(container, v.Value, group, isRequestType)
		case spec.ArrayType:
			joinContainer(container, v.Value, group, isRequestType)
		case spec.DefineStruct:
			joinContainer(container, m.Type, group, isRequestType)
		}
	}
	if typeName == "" {
		return
	}
	_, ok = container[typeName]
	if !ok {
		container[typeName] = make(map[string]int, 0)
	}
	container[typeName][group] = 1
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
