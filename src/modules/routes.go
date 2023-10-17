package modules

import (
	"vestislabs/http2-uploader/src/modules/download"
	"vestislabs/http2-uploader/src/modules/uploader"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo) {
	v1 := app.Group("/v1")

	uploader.RouterV1(v1.Group("/uploader"))
	download.RouterV1(v1.Group("/download"))

	v1Auth := app.Group("/v1")
	v1Auth.Use(TokenValidation)
}
