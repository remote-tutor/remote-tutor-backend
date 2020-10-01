package files

import (
	"github.com/labstack/echo"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

func ReadFromSource(c echo.Context, formFileName string) (string, multipart.File, error) {
	file, err := c.FormFile(formFileName)
	if err != nil {
		return "", nil, err
	}
	src, err := file.Open()
	if err != nil {
		return "", nil, err
	}
	defer src.Close()
	return file.Filename, src, nil
}

func CopyFileToDestination(fullFileName string, src multipart.File) error {
	// add file to destination
	dst, err := os.Create(fullFileName)
	if err != nil {
		return err
	}
	// copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func CreateDirectoryIfNotExist(directoryName string) {
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		os.Mkdir(directoryName, os.FileMode(int(0777)))
	}
}

func DeleteDirectory(directoryName string) {
	os.RemoveAll(directoryName)
}

func GetFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
