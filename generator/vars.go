package generator

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	etcDir = "etc"

	typesPacket = "types"
	typesDir    = typesPacket

	configPacket = "config"
	configDir    = configPacket

	handlerPacket = "handler"
	handlerDir    = handlerPacket

	routesPacket = "routes"
	routesDir    = routesPacket

	logicPacket = "logic"
	logicDir    = logicPacket

	middlewarePacket = "middleware"
	middlewareDir    = middlewarePacket

	typePkgAlias = "Type"

	groupProperty = "group"

	dirFmt = "go/zero"
)

var title = cases.Title(language.English)
