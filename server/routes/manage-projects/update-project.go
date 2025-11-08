package manageprojects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dbconfig"
	"strings"
	"time"
)


func UpdateProjectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST or PUT."}`, http.StatusMethodNotAllowed)
		return
	}


	var input struct {
		ProjectID          int     `json:"project_id"`
		FDNumber           *string `json:"fd_number,omitempty"`
		ProjectName        *string `json:"project_name,omitempty"`
		ProjectDescription *string `json:"project_description,omitempty"`
		ProjectBudget      *int    `json:"project_amount,omitempty"`
		Status             *string `json:"status,omitempty"`
		AgreementDocs      *string `json:"agreement_docs,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.ProjectID <= 0 {
		http.Error(w, `{"success": false, "message": "project_id is required and must be valid"}`, http.StatusBadRequest)
		return
	}

	setClauses := []string{}
	params := []interface{}{}
	paramIndex := 1

	if input.FDNumber != nil {
		setClauses = append(setClauses, fmt.Sprintf("fd_number = $%d", paramIndex))
		params = append(params, *input.FDNumber)
		paramIndex++
	}
	if input.ProjectName != nil {
		setClauses = append(setClauses, fmt.Sprintf("project_name = $%d", paramIndex))
		params = append(params, *input.ProjectName)
		paramIndex++
	}
	if input.ProjectDescription != nil {
		setClauses = append(setClauses, fmt.Sprintf("project_description = $%d", paramIndex))
		params = append(params, *input.ProjectDescription)
		paramIndex++
	}
	if input.ProjectBudget != nil {
		setClauses = append(setClauses, fmt.Sprintf("project_amount = $%d", paramIndex))
		params = append(params, *input.ProjectBudget)
		paramIndex++
	}
	if input.Status != nil {
		setClauses = append(setClauses, fmt.Sprintf("status = $%d", paramIndex))
		params = append(params, *input.Status)
		paramIndex++
	}
	if input.AgreementDocs != nil {
		setClauses = append(setClauses, fmt.Sprintf("agreement_docs = $%d", paramIndex))
		params = append(params, *input.AgreementDocs)
		paramIndex++
	}

	if len(setClauses) == 0 {
		http.Error(w, `{"success": false, "message": "No fields provided for update"}`, http.StatusBadRequest)
		return
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", paramIndex))
	params = append(params, time.Now())
	paramIndex++


	query := fmt.Sprintf(`UPDATE manage_project SET %s WHERE project_id = $%d`, strings.Join(setClauses, ", "), paramIndex)
	params = append(params, input.ProjectID)

	result, err := dbconfig.Database.Exec(query, params...)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database update error", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error fetching update result"}`, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No project found with the given project_id"}`, http.StatusNotFound)
		return
	}


	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    true,
		"message":    "Project updated successfully",
		"project_id": input.ProjectID,
	})
}
