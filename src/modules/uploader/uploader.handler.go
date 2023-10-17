package uploader

import (
	"net/http"

	"vestislabs/http2-uploader/src"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

// @Summary Upload and Convert Video
// @Description Uploads a video, converts it, and then sends back the converted video.
// @Tags Video
// @Accept mpfd
// @Produce video/mp4
// @Param file formData file true "Video file"
// @Success 200 {file} byte "Converted video file"
// @Failure 400 {object} src.JSONResponse "Bad request"
// @Failure 500 {object} src.JSONResponse "Internal server error"
// @Router /v1/uploader/ [post]
func VideoHandler(c echo.Context) error {
	// Dapatkan file dari form multipart
	file, err := c.FormFile("file")
	if err != nil {
		return src.SendJSONResponse(c, http.StatusBadRequest, false, "Failed to get the uploaded file.")
	}

	file_src, err := file.Open()
	if err != nil {
		return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Failed to open the uploaded file.")
	}
	defer file_src.Close()

	// Pastikan ukuran file tidak lebih dari 200MB
	if file.Size > src.LimitUploadSize {
		return src.SendJSONResponse(c, http.StatusBadRequest, false, "File size exceeds the limit.")
	}

	// Memeriksa apakah file adalah video
	buf := make([]byte, 261)
	_, err = file_src.Read(buf)
	if err != nil {
		return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Failed to read file.")
	}
	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown || !filetype.IsVideo(buf) {
		return src.SendJSONResponse(c, http.StatusBadRequest, false, "File is not a valid video.")
	}
	file_src.Seek(0, 0) // Kembali ke awal file setelah membaca

	return VideoService(c, file, file_src)
}
