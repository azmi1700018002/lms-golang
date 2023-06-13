package helper

import (
	"fmt"
	"io"

	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFile(file *multipart.FileHeader, folderPath string) (string, error) {
	// Generate random file name to avoid naming conflicts
	fileName := GenerateRandomFileName(file.Filename)

	// Create the destination folder if it doesn't exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", err
		}
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create the file at the destination folder
	dstPath := filepath.Join(folderPath, fileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the content of the file to the destination file
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return dstPath, nil
}

// GenerateRandomFileName generates a random file name
func GenerateRandomFileName(originalFileName string) string {
	extension := filepath.Ext(originalFileName)
	name := originalFileName[0 : len(originalFileName)-len(extension)]
	randomName := RandomString(10)
	return fmt.Sprintf("%s_%s%s", name, randomName, extension)
}

// RandomString generates a random string of a given length
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
