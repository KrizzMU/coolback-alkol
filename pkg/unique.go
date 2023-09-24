package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetPath(name string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName)

	_, err := os.Stat(dirPath)
	if !os.IsNotExist(err) {
		return dirPath, nil
	}

	return "", fmt.Errorf("folder with that name doesn't exist")
}

func CreateUniqueFolder(name string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName)

	splitFolderName := strings.Split(dirPath, "\\")
	dbFolderName := splitFolderName[len(splitFolderName)-1]

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		return dbFolderName, nil
	}

	return "", fmt.Errorf("directory already exists")
}

func CreateUniqueFile(file multipart.File, fileName string, name string, folder string, requiredExt string) (string, error) {
	defer file.Close()

	ext := filepath.Ext(fileName)
	if ext != requiredExt {
		return "", fmt.Errorf("wrong file extension")
	}

	nameWithoutExt := name

	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(nameWithoutExt, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	filePath := filepath.Join(folder, uniqueName+ext)

	splitFilePath := strings.Split(filePath, "\\")
	dbFileName := splitFilePath[len(splitFilePath)-1]

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		out, err := os.Create(filePath)
		if err != nil {
			return "", err
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			return "", err
		}

		if err = out.Close(); err != nil {
			return "", err
		}

		return dbFileName, nil
	}

	if err = file.Close(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("file already exists")
}
