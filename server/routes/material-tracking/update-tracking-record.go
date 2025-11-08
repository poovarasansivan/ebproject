package materialtracking

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dbconfig"
	"strings"
	"time"
)

func UpdateTrackingRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST or PUT."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		TrackingID     int     `json:"tracking_id"`
		StockID        *int    `json:"stock_id,omitempty"`
		ProjectID      *int    `json:"project_id,omitempty"`
		NoOfStocks     *int    `json:"no_of_stocks_used,omitempty"`
		ImageProof     *string `json:"image_proof,omitempty"`
		MachineRunTime *string `json:"machines_running_time,omitempty"`
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

	setClauses := []string{}
	params := []interface{}{}
	paramIndex := 1

	if input.StockID != nil {
		setClauses = append(setClauses, fmt.Sprintf("stock_id = $%d", paramIndex))
		params = append(params, *input.StockID)
		paramIndex++
	}
	if input.ProjectID != nil {
		setClauses = append(setClauses, fmt.Sprintf("project_id = $%d", paramIndex))
		params = append(params, *input.ProjectID)
		paramIndex++
	}
	if input.NoOfStocks != nil {
		setClauses = append(setClauses, fmt.Sprintf("no_of_stocks_used = $%d", paramIndex))
		params = append(params, *input.NoOfStocks)
		paramIndex++
	}
	if input.ImageProof != nil {
		setClauses = append(setClauses, fmt.Sprintf("image_proof = $%d", paramIndex))
		params = append(params, *input.ImageProof)
		paramIndex++
	}
	if input.MachineRunTime != nil {
		setClauses = append(setClauses, fmt.Sprintf("machines_running_time = $%d", paramIndex))
		params = append(params, *input.MachineRunTime)
		paramIndex++
	}

	if len(setClauses) == 0 {
		http.Error(w, `{"success": false, "message": "No fields provided to update"}`, http.StatusBadRequest)
		return
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", paramIndex))
	params = append(params, time.Now())
	paramIndex++

	query := fmt.Sprintf(`
		UPDATE material_tracking
		SET %s
		WHERE tracking_id = $%d
	`, strings.Join(setClauses, ", "), paramIndex)

	params = append(params, input.TrackingID)

	result, err := dbconfig.Database.Exec(query, params...)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database update error", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error checking update result"}`, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No record found for the given tracking_id"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"message":     "Material tracking record updated successfully",
		"tracking_id": input.TrackingID,
		"updated_at":  time.Now(),
	})
}
