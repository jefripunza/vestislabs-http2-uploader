package download

import "github.com/labstack/echo/v4"

func RouterV1(app *echo.Group) {
	app.GET("/:filename", DownloadHandler)
}
