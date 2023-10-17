package src

import "os"

func MakeDir() {
	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		os.Mkdir(UploadPath, 0755)
	}
}
