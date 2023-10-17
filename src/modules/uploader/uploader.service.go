package uploader

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"vestislabs/http2-uploader/src"

	"github.com/labstack/echo/v4"
)

func VideoService(c echo.Context, file *multipart.FileHeader, file_src multipart.File) error {
	// Simpan file
	dst, err := os.Create(filepath.Join(src.UploadPath, file.Filename))
	if err != nil {
		return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Failed to save the uploaded file.")
	}

	if _, err = io.Copy(dst, file_src); err != nil {
		dst.Close()
		return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Failed to copy the uploaded file.")
	}

	// Pastikan menutup file setelah penyalinan selesai
	dst.Close()

	// Konversi menggunakan FFmpeg
	inputFilePath := filepath.Join(src.UploadPath, file.Filename)
	outputFilePath := filepath.Join(src.UploadPath, "converted_"+file.Filename)

	cmd := exec.Command("ffmpeg", "-i", inputFilePath, outputFilePath)

	errChannel := make(chan error)

	go func() {
		_, err := cmd.Output()
		errChannel <- err
	}()

	select {
	case err := <-errChannel:
		if err != nil {
			fmt.Println(err)
			return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Failed to convert video.")
		}
	case <-time.After(10 * time.Minute): // Timeout set to 10 minutes
		return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Video conversion timed out.")
	}

	fileToStream, err := os.Open(outputFilePath)
	if err != nil {
		fmt.Println(err)
		return src.SendJSONResponse(c, http.StatusInternalServerError, false, "Failed to open the converted file for streaming.")
	}
	defer fileToStream.Close()

	/*
		TODO: apapun yang terjadi, akan jadi mp4 saat stream
	*/
	// c.Response().Header().Set(echo.HeaderContentType, "video/mp4")
	// c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "converted_"+file.Filename))
	// if err := c.Stream(http.StatusOK, "video/mp4", fileToStream); err != nil {
	// 	return err
	// }

	contentType := src.GetContentTypeByExtension(file.Filename)
	c.Response().Header().Set(echo.HeaderContentType, contentType)
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "converted_"+file.Filename))
	if err := c.Stream(http.StatusOK, contentType, fileToStream); err != nil {
		fmt.Println(err)
		return err
	}

	// Hapus file setelah selesai dikirim
	if err := os.Remove(outputFilePath); err != nil {
		fmt.Println(err)
		fmt.Println("Failed to delete the file after streaming:", err)
	}

	return nil
}
