package materialtracking

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetTrackingRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use GET."}`, http.StatusMethodNotAllowed)
		return
	}

	query := `
		SELECT 
			mt.tracking_id,
			mt.stock_id,
			ms.stock_name,
			mt.project_id,
			mp.project_name,
			mt.no_of_stocks_used,
			mt.image_proof,
			mt.machines_running_time,
			mt.created_at,
			mt.updated_at
		FROM material_tracking mt
		LEFT JOIN manage_stocks ms ON mt.stock_id = ms.stock_id
		LEFT JOIN manage_project mp ON mt.project_id = mp.project_id
		ORDER BY mt.created_at DESC;
	`

	rows, err := dbconfig.Database.Query(query)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database query error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var records []models.TrackingRecordResponse

	for rows.Next() {
		var record models.TrackingRecordResponse
		err := rows.Scan(
			&record.TrackingID,
			&record.StockID,
			&record.StockName,
			&record.ProjectID,
			&record.ProjectName,
			&record.NoOfStocks,
			&record.ImageProof,
			&record.MachineRunTime,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			http.Error(w, `{"success": false, "message": "Error scanning record: `+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
		records = append(records, record)
	}

	if len(records) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "No tracking records found",
			"data":    []models.TrackingRecordResponse{},
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Tracking records fetched successfully",
		"data":    records,
	})
}
