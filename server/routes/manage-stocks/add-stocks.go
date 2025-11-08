package managestocks

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
	"time"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST."}`, http.StatusMethodNotAllowed)
		return
	}

	var input models.ManageStocksModel
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.StockName == "" || input.StockType == "" {
		http.Error(w, `{"success": false, "message": "stock_name and stock_type are required"}`, http.StatusBadRequest)
		return
	}

	if input.AvailableStock < 0 {
		input.AvailableStock = 0
	}
	if input.StockConsumed < 0 {
		input.StockConsumed = 0
	}
	if input.ReturnedStock < 0 {
		input.ReturnedStock = 0
	}

	query := `
		INSERT INTO manage_stocks (stock_name, stock_type, available_stock, stock_delivered_out, returned_stock, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING stock_id
	`

	var newStockID int
	err := dbconfig.Database.QueryRow(
		query,
		input.StockName,
		input.StockType,
		input.AvailableStock,
		input.StockConsumed,
		input.ReturnedStock,
		time.Now(),
		time.Now(),
	).Scan(&newStockID)

	if err != nil {
		http.Error(w, `{"success": false, "message": "Database insertion failed", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Stock added successfully",
		"stock_id": newStockID,
	})
}
