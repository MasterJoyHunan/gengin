package tpl

import _ "embed"

var (
	//go:embed etc.tpl
	EtcTemplate string

	//go:embed config.tpl
	ConfigTemplate string

	//go:embed svc_context.tpl
	SvcContextTemplate string

	//go:embed main.tpl
	MainTemplate string

	//go:embed middleware.tpl
	MiddlewareTemplate string

	//go:embed types.tpl
	TypesTemplate string

	//go:embed logic.tpl
	LogicTemplate string

	//go:embed routes.tpl
	RoutesTemplate string

	//go:embed routes_setup.tpl
	RoutesSetupTemplate string

	//go:embed handler.tpl
	HandlerTemplate string

	//go:embed response.tpl
	ResponseTemplate string

	//go:embed i18n.tpl
	I18nTemplate string
)
