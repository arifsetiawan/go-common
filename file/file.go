package file

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
)

// SaveFileHeader will file content to file
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

// GetFileHeaderContent will return file content
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

// ReadFileHeader will return file content, hash, size
func ReadFileHeader(fileHeader *multipart.FileHeader) (io.Reader, string, int64, error) {
	// Open file
	fileSrc, err := fileHeader.Open()
	if err != nil {
		return nil, "", 0, err
	}
	defer fileSrc.Close()

	var buf bytes.Buffer
	hash := sha1.New()
	n, err := io.Copy(&buf, io.TeeReader(fileSrc, hash))

	if err != nil {
		return nil, "", 0, err
	}

	return &buf, hex.EncodeToString(hash.Sum(nil)), n, nil
}
