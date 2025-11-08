package managestocks

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
)

func DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		StockID int `json:"stock_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.StockID == 0 {
		http.Error(w, "Missing or invalid stock_id", http.StatusBadRequest)
		return
	}

	db := dbconfig.Database 

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM manage_stocks WHERE stock_id=$1)`
	err := db.QueryRow(checkQuery, req.StockID).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error while checking stock", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	deleteQuery := `DELETE FROM manage_stocks WHERE stock_id=$1`
	result, err := db.Exec(deleteQuery, req.StockID)
	if err != nil {
		http.Error(w, "Failed to delete stock", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No record deleted", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"message": "Stock deleted successfully",
		"stockid": req.StockID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
