package advanced

import (
	_ "embed"
)

//go:embed static/htmx/hello.templ.tmpl
var helloTemplate []byte

//go:embed static/htmx/base.templ.tmpl
var baseTemplTemplate []byte

//go:embed static/htmx/htmx.min.js.tmpl
var htmxMinJsTemplate []byte

//go:embed static/htmx/efs.go.tmpl
var efsTemplate []byte

//go:embed static/htmx/hello.go.tmpl
var helloGoTemplate []byte

//go:embed static/htmx/hello_fiber.go.tmpl
var helloFiberGoTemplate []byte

//go:embed static/htmx/routes/http_router.tmpl
var httpRouterHtmxTemplRoutes []byte

//go:embed static/htmx/routes/standard_library.tmpl
var stdLibHtmxTemplRoutes []byte

//go:embed static/htmx/imports/standard_library.tmpl
var stdLibHtmxTemplImports []byte

//go:embed static/htmx/routes/chi.tmpl
var chiHtmxTemplRoutes []byte

//go:embed static/htmx/routes/gin.tmpl
var ginHtmxTemplRoutes []byte

//go:embed static/htmx/routes/gorilla.tmpl
var gorillaHtmxTemplRoutes []byte

//go:embed static/htmx/routes/echo.tmpl
var echoHtmxTemplRoutes []byte

//go:embed static/htmx/routes/fiber.tmpl
var fiberHtmxTemplRoutes []byte

//go:embed static/htmx/imports/fiber.tmpl
var fiberHtmxTemplImports []byte

func EchoHtmxTemplRoutesTemplate() []byte {
	return echoHtmxTemplRoutes
}

func GorillaHtmxTemplRoutesTemplate() []byte {
	return gorillaHtmxTemplRoutes
}

func ChiHtmxTemplRoutesTemplate() []byte {
	return chiHtmxTemplRoutes
}

func GinHtmxTemplRoutesTemplate() []byte {
	return ginHtmxTemplRoutes
}

func HttpRouterHtmxTemplRoutesTemplate() []byte {
	return httpRouterHtmxTemplRoutes
}

func StdLibHtmxTemplRoutesTemplate() []byte {
	return stdLibHtmxTemplRoutes
}

func StdLibHtmxTemplImportsTemplate() []byte {
	return stdLibHtmxTemplImports
}

func HelloTemplate() []byte {
	return helloTemplate
}

func BaseTemplate() []byte {
	return baseTemplTemplate
}

func HtmxJSTemplate() []byte {
	return htmxMinJsTemplate
}

func EfsTemplate() []byte {
	return efsTemplate
}

func HelloGoTemplate() []byte {
	return helloGoTemplate
}

func HelloFiberGoTemplate() []byte {
	return helloFiberGoTemplate
}

func FiberHtmxTemplRoutesTemplate() []byte {
	return fiberHtmxTemplRoutes
}

func FiberHtmxTemplImportsTemplate() []byte {
	return fiberHtmxTemplImports
}
