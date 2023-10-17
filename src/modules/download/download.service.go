package download

import (
	"net/http"
	"os"
	"path/filepath"
	"vestislabs/http2-uploader/src"

	"github.com/labstack/echo/v4"
)

func DownloadService(c echo.Context, filename string) error {
	// Pastikan file ada
	filepath := filepath.Join(src.UploadPath, filename)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return src.SendJSONResponse(c, http.StatusNotFound, false, "File not found.")
	}

	return c.Attachment(filepath, filename)
}
