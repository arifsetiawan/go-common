package file

import (
	"os"
	"string"
	"bytes"
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

func GetFileHeaderContent(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	// Copy
	if _, err = io.Copy(buf, src); err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}