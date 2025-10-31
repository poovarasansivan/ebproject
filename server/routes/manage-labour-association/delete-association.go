package managelabourassociation

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
)

func DeleteAssociationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method. Only DELETE is allowed.", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		AssociationID int `json:"association_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.AssociationID == 0 {
		http.Error(w, "Missing or invalid association_id", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM manager_labour_association WHERE association_id = $1`

	result, err := dbconfig.Database.Exec(query, requestData.AssociationID)
	if err != nil {
		http.Error(w, "Database delete error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "No association found with the provided ID",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Association deleted successfully",
	})
}
