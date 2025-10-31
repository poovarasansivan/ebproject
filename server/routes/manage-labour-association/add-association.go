package managelabourassociation

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"time"
)

func AddManagerLabourAssociationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Only POST allowed."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		ManagerID string `json:"manager_id"`
		LabourID  string `json:"labour_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.ManagerID == "" || input.LabourID == "" {
		http.Error(w, `{"success": false, "message": "Both manager_id and labour_id are required"}`, http.StatusBadRequest)
		return
	}

	currentTime := time.Now()

	query := `
		INSERT INTO manager_labour_association (manager_id, labour_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING association_id
	`

	var insertedID int
	err := dbconfig.Database.QueryRow(query, input.ManagerID, input.LabourID, currentTime, currentTime).Scan(&insertedID)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database insertion error", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Manager-Labour association added successfully",
		"data": map[string]interface{}{
			"association_id": insertedID,
			"manager_id":     input.ManagerID,
			"labour_id":      input.LabourID,
			"created_at":     currentTime,
			"updated_at":     currentTime,
		},
	})
}
