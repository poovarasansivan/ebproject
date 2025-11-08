package handlefile

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(r *http.Request, formKey string) (string, error) {
	file, handler, err := r.FormFile(formKey)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}

	projectDocsDir := filepath.Join(wd, "project_docs")
	if _, err := os.Stat(projectDocsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(projectDocsDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
	filePath := filepath.Join(projectDocsDir, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return fileName, nil
}
