package manageprojects

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetProjectDetails(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Use GET.", http.StatusMethodNotAllowed)
		return
	}

	query := `
		SELECT 
			project_id,
			fd_number,
			project_name,
			project_description,
			project_amount,
			status,
			agreement_docs,
			created_at,
			updated_at
		FROM manage_project
		ORDER BY created_at DESC;
	`

	rows, err := dbconfig.Database.Query(query)
	if err != nil {
		http.Error(w, "Database query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var projects []models.ProjectModel

	for rows.Next() {
		var p models.ProjectModel
		err := rows.Scan(
			&p.ProjectID,
			&p.FDNumber,
			&p.ProjectName,
			&p.ProjectDescription,
			&p.ProjectBudget,
			&p.Status,
			&p.AgreementDocs,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning project data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		projects = append(projects, p)
	}

	if len(projects) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "No projects found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"count":    len(projects),
		"projects": projects,
	})
}
