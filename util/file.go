package util

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/xid"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getWorkingDirectory(path string, fileName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(wd, "assets", path, fileName)

	return fullPath, nil
}

func UploadFile(path string, graphqlUpload graphql.Upload) (string, error) {
	megaByte := 1
	if graphqlUpload.Size > (int64(megaByte) * 1024 * 1024) {
		return "", fmt.Errorf("file too large")
	}

	split := strings.Split(graphqlUpload.Filename, ".")
	extension := fmt.Sprintf(".%s", split[len(split)-1])

	guid := xid.New()
	fileName := fmt.Sprintf("%s%s", guid.String(), extension)

	fullPath, err := getWorkingDirectory(path, fileName)
	if err != nil {
		return "", err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, graphqlUpload.File)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func DeleteFile(path string, fileName string) error {
	fullPath, err := getWorkingDirectory(path, fileName)
	if err != nil {
		return err
	}

	err = os.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}
