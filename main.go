package main

import (
	"os"
	"vestislabs/http2-uploader/src"
	"vestislabs/http2-uploader/src/modules"

	"github.com/labstack/echo/v4"

	_ "vestislabs/http2-uploader/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @Title Vestislabs HTTP2 Uploader
// @Version 1.0
// @Description This is an API for example create stream video after upload & convert
// @TermsOfService Vestislabs
// @Contact.name  Jefri Herdi Triyanto
// @Contact.url   https://github.com/jefripunza/vestislabs-http2-uploader
// @Contact.email jefriherditriyanto@gmail.com
func main() {
	src.MakeDir()

	app := echo.New()

	app.GET("/swagger/*", echoSwagger.WrapHandler)
	modules.GlobalMiddlewares(app)
	modules.Routes(app)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "1337"
	}

	app.Logger.Fatal(app.Start(":" + httpPort))
}
