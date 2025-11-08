package handletrackingproofs

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

func UploadTrackingProof(r *http.Request, formKey string) (string, error) {
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

	uploadDir := filepath.Join(wd, "material_tracking_proofs")
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
	fileExt := strings.ToLower(filepath.Ext(handler.Filename))
	allowedExtensions := []string{".pdf", ".png", ".jpg", ".jpeg"}
	for _, ext := range allowedExtensions {
		if fileExt == ext {
			return true
		}
	}
	return false
}

func sanitizeFileName(name string) string {
	name = filepath.Base(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "..", "_")
	name = strings.ReplaceAll(name, "/", "_")
	return name
}
