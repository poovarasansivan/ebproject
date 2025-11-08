package materialtracking

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/handletrackingproofs"
	"server/models"
	"strconv"
	"time"
)

func AddTrackingRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST."}`, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, `{"success": false, "message": "Error parsing form data: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	var record models.MaterialTrackingModel

	record.StockID, _ = strconv.Atoi(r.FormValue("stock_id"))
	record.ProjectID, _ = strconv.Atoi(r.FormValue("project_id"))
	record.NoOfStocks, _ = strconv.Atoi(r.FormValue("no_of_stocks_used"))
	record.MachineRunTime = r.FormValue("machines_running_time")

	fileName, err := handletrackingproofs.UploadTrackingProof(r, "image_proof")
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error uploading proof file: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	record.ImageProof = fileName

	now := time.Now().Format("2006-01-02 15:04:05")

	query := `
		INSERT INTO material_tracking 
		(stock_id, project_id, no_of_stocks_used, image_proof, machines_running_time, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING tracking_id
	`

	err = dbconfig.Database.QueryRow(
		query,
		record.StockID,
		record.ProjectID,
		record.NoOfStocks,
		record.ImageProof,
		record.MachineRunTime,
		now,
		now,
	).Scan(&record.TrackingID)

	if err != nil {
		http.Error(w, `{"success": false, "message": "Database insertion error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"message":     "Material tracking record added successfully",
		"tracking_id": record.TrackingID,
	})
}
