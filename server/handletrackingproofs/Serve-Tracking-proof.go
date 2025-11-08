package handletrackingproofs

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)
  
func ServerTrackingProof(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	
	if fileName == "" {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}
	if strings.Contains(fileName, "..") {
		http.Error(w, "Invalid file name", http.StatusBadRequest)
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Failed to get working directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(wd, "material_tracking_proofs", fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	ext := strings.ToLower(filepath.Ext(fileName))
	allowedExts := []string{".pdf", ".png", ".jpg", ".jpeg"}

	isAllowed := false
	for _, allowed := range allowedExts {
		if ext == allowed {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		http.Error(w, "Unsupported file type", http.StatusForbidden)
		return
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		switch ext {
		case ".pdf":
			mimeType = "application/pdf"
		case ".jpg", ".jpeg":
			mimeType = "image/jpeg"
		case ".png":
			mimeType = "image/png"
		default:
			mimeType = "application/octet-stream"
		}
	}

	w.Header().Set("Content-Type", mimeType)
	http.ServeFile(w, r, filePath)
}