package file

import (
	"os"
	"io"
	"mime/multipart"
)

func SaveFileHeader(file *multipart.FileHeader, filename string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}