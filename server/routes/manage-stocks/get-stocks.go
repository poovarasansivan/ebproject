package managestocks

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	db := dbconfig.Database

	query := `
		SELECT 
			stock_id, 
			stock_name, 
			stock_type, 
			available_stock, 
			stock_delivered_out, 
			returned_stock, 
			created_at, 
			updated_at
		FROM manage_stocks
		ORDER BY stock_id ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Database query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var stocks []models.ManageStocksModel

	for rows.Next() {
		var s models.ManageStocksModel
		err := rows.Scan(
			&s.StockID,
			&s.StockName,
			&s.StockType,
			&s.AvailableStock,
			&s.StockConsumed,
			&s.ReturnedStock,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning rows: "+err.Error(), http.StatusInternalServerError)
			return
		}
		stocks = append(stocks, s)
	}

	if len(stocks) == 0 {
		http.Error(w, "No stocks found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}
