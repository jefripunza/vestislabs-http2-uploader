package modules

import (
	"net/http"

	"vestislabs/http2-uploader/src"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GlobalMiddlewares(app *echo.Echo) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

func TokenValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		Token := c.Request().Header.Get("Authorization")
		if Token == "" {
			return src.SendJSONResponse(c, http.StatusUnauthorized, false, "No token provided.")
		}
		return next(c)
	}
}
