package download

import (
	"github.com/labstack/echo/v4"
)

// @Summary Download a file
// @Description Download a file by filename
// @Produce  octet-stream
// @Param filename path string true "Filename"
// @Success 200 {file} byte
// @Failure 404 {object} src.JSONResponse
// @Router /v1/download/{filename}/ [get]
func DownloadHandler(c echo.Context) error {
	filename := c.Param("filename")

	return DownloadService(c, filename)
}
