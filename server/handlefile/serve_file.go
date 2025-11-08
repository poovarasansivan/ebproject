package handlefile

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServePDFHandler(w http.ResponseWriter, r *http.Request) {
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

	filePath := filepath.Join(wd, "project_docs", fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if !strings.HasSuffix(strings.ToLower(fileName), ".pdf") {
		http.Error(w, "Only PDF files can be served", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, filePath)
}
