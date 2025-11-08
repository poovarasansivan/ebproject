package materialtracking

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"server/dbconfig"
	"time"
)

func DeleteTrackingRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete && r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use DELETE or POST."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		TrackingID int `json:"tracking_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.TrackingID <= 0 {
		http.Error(w, `{"success": false, "message": "Valid tracking_id is required"}`, http.StatusBadRequest)
		return
	}

	var imageProof sql.NullString
	err := dbconfig.Database.QueryRow(`
		SELECT image_proof 
		FROM material_tracking 
		WHERE tracking_id = $1
	`, input.TrackingID).Scan(&imageProof)

	if err == sql.ErrNoRows {
		http.Error(w, `{"success": false, "message": "No record found with the given tracking_id"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"success": false, "message": "Database query error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	deleteQuery := `DELETE FROM material_tracking WHERE tracking_id = $1`
	result, err := dbconfig.Database.Exec(deleteQuery, input.TrackingID)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database delete error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error checking deletion result"}`, http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No tracking record deleted. Please check tracking_id."}`, http.StatusNotFound)
		return
	}

	if imageProof.Valid && imageProof.String != "" {
		filePath := filepath.Join("server", "material_tracking_proofs", imageProof.String)
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err != nil {
				http.Error(w, `{"success": true, "message": "Record deleted but failed to delete proof file: `+err.Error()+`"}`, http.StatusInternalServerError)
				return
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":       true,
		"message":       "Material tracking record deleted successfully",
		"tracking_id":   input.TrackingID,
		"proof_deleted": imageProof.Valid && imageProof.String != "",
		"deleted_at":    time.Now(),
	})
}
