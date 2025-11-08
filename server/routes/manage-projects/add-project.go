package manageprojects

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/handlefile"
)

func AddNewProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST."}`, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, `{"success": false, "message": "Error parsing form data: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	fdNumber := r.FormValue("fd_number")
	projectName := r.FormValue("project_name")
	projectDescription := r.FormValue("project_description")
	projectBudget := r.FormValue("project_amount")
	status := r.FormValue("status")

	if fdNumber == "" || projectName == "" {
		http.Error(w, `{"success": false, "message": "fd_number and project_name are required"}`, http.StatusBadRequest)
		return
	}

	fileName, err := handlefile.UploadFile(r, "agreement_docs")
	if err != nil {
		http.Error(w, `{"success": false, "message": "File upload failed: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	query := `
		INSERT INTO manage_project 
		(fd_number, project_name, project_description, project_amount, status, agreement_docs, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING project_id;
	`

	var projectID int
	err = dbconfig.Database.QueryRow(query,
		fdNumber,
		projectName,
		projectDescription,
		projectBudget,
		status,
		fileName, 
	).Scan(&projectID)

	if err != nil {
		http.Error(w, `{"success": false, "message": "Database insert error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":        true,
		"message":        "Project created successfully",
		"project_id":     projectID,
		"agreement_docs": fileName,
	})
}
