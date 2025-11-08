package manageprojects

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"server/dbconfig"
	"time"
)

func DeleteProjectDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete && r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use DELETE or POST."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		ProjectID int `json:"project_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.ProjectID <= 0 {
		http.Error(w, `{"success": false, "message": "Valid project_id is required"}`, http.StatusBadRequest)
		return
	}

	var docFile sql.NullString
	err := dbconfig.Database.QueryRow(`
		SELECT agreement_docs 
		FROM manage_project 
		WHERE project_id = $1
	`, input.ProjectID).Scan(&docFile)

	if err == sql.ErrNoRows {
		http.Error(w, `{"success": false, "message": "No project found with the given project_id"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"success": false, "message": "Database error while fetching project details: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}


	deleteQuery := `DELETE FROM manage_project WHERE project_id = $1`
	result, err := dbconfig.Database.Exec(deleteQuery, input.ProjectID)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Failed to delete project", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No project deleted. Please check project_id."}`, http.StatusNotFound)
		return
	}

	if docFile.Valid && docFile.String != "" {
		wd, _ := os.Getwd()

		filePath := filepath.Join(wd, "project_docs", filepath.Base(docFile.String))

		if _, err := os.Stat(filePath); err == nil {
			if rmErr := os.Remove(filePath); rmErr == nil {
			fmt.Print("Project document deleted")
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":       true,
		"message":       "Project deleted successfully",
		"project_id":    input.ProjectID,
		"deleted_at":    time.Now(),
		"doc_deleted":   docFile.Valid && docFile.String != "",
	})
}
