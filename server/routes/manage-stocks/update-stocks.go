package managestocks

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"strconv"
	"strings"
	"time"
)


func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut && r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use PUT or POST."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		StockID        int     `json:"stock_id"`
		StockName      *string `json:"stock_name,omitempty"`
		StockType      *string `json:"stock_type,omitempty"`
		AvailableStock *int    `json:"available_stock,omitempty"`
		StockConsumed  *int    `json:"stock_delivered_out,omitempty"`
		ReturnedStock  *int    `json:"returned_stock,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.StockID <= 0 {
		http.Error(w, `{"success": false, "message": "stock_id is required and must be valid"}`, http.StatusBadRequest)
		return
	}

	var setClauses []string
	var params []interface{}
	paramIndex := 1

	if input.StockName != nil {
		setClauses = append(setClauses, "stock_name = $"+itoa(paramIndex))
		params = append(params, *input.StockName)
		paramIndex++
	}
	if input.StockType != nil {
		setClauses = append(setClauses, "stock_type = $"+itoa(paramIndex))
		params = append(params, *input.StockType)
		paramIndex++
	}
	if input.AvailableStock != nil {
		setClauses = append(setClauses, "available_stock = $"+itoa(paramIndex))
		params = append(params, *input.AvailableStock)
		paramIndex++
	}
	if input.StockConsumed != nil {
		setClauses = append(setClauses, "stock_delivered_out = $"+itoa(paramIndex))
		params = append(params, *input.StockConsumed)
		paramIndex++
	}
	if input.ReturnedStock != nil {
		setClauses = append(setClauses, "returned_stock = $"+itoa(paramIndex))
		params = append(params, *input.ReturnedStock)
		paramIndex++
	}

	if len(setClauses) == 0 {
		http.Error(w, `{"success": false, "message": "No fields provided for update"}`, http.StatusBadRequest)
		return
	}

	setClauses = append(setClauses, "updated_at = $"+itoa(paramIndex))
	params = append(params, time.Now())
	paramIndex++

	query := `UPDATE manage_stocks SET ` + strings.Join(setClauses, ", ") + ` WHERE stock_id = $` + itoa(paramIndex)
	params = append(params, input.StockID)

	result, err := dbconfig.Database.Exec(query, params...)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database update failed", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error checking update result"}`, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No record found with the given stock_id"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Stock updated successfully",
		"stock_id": input.StockID,
	})
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
