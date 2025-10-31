package managelabourassociation

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
	"time"
)

func UpdateAssociationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method. Only PUT is allowed.", http.StatusMethodNotAllowed)
		return
	}

	var association models.LabourAssociationModel
	if err := json.NewDecoder(r.Body).Decode(&association); err != nil {
		http.Error(w, "Invalid JSON input: "+err.Error(), http.StatusBadRequest)
		return
	}

	if association.AssociationID == 0 {
		http.Error(w, "Missing association_id", http.StatusBadRequest)
		return
	}
	if association.ManagerID == "" || association.LabourID == "" {
		http.Error(w, "manager_id and labour_id are required", http.StatusBadRequest)
		return
	}

	query := `
		UPDATE manager_labour_association 
		SET manager_id = $1, labour_id = $2, updated_at = $3
		WHERE association_id = $4
	`
	_, err := dbconfig.Database.Exec(query,
		association.ManagerID,
		association.LabourID,
		time.Now(),
		association.AssociationID,
	)

	if err != nil {
		http.Error(w, "Database update error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Labour association updated successfully",
		"data":    association,
	})
}
