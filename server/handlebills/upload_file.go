package handlebills

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)


func UploadExpenseFile(r *http.Request, formKey string) (string, error) {
	file, handler, err := r.FormFile(formKey)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	if !isAllowedFileType(handler) {
		return "", fmt.Errorf("invalid file type: only PDF and image files (PNG/JPG/JPEG) are allowed")
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %v", err)
	}

	uploadDir := filepath.Join(wd, "expense_bills")

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), sanitizeFileName(handler.Filename))
	filePath := filepath.Join(uploadDir, fileName)

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

func isAllowedFileType(handler *multipart.FileHeader) bool {
	allowedExts := []string{".pdf", ".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(handler.Filename))

	for _, allowed := range allowedExts {
		if ext == allowed {
			return true
		}
	}
	return false
}

func sanitizeFileName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "..", "_")
	return filepath.Base(name)
}
