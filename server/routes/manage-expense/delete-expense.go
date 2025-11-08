package manageexpense

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"server/dbconfig"
	"time"
)

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use DELETE."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		ExpenseID int `json:"expense_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.ExpenseID <= 0 {
		http.Error(w, `{"success": false, "message": "Valid expense_id is required"}`, http.StatusBadRequest)
		return
	}

	var billFile sql.NullString
	err := dbconfig.Database.QueryRow(`
		SELECT bill_docs FROM manage_expenditure WHERE expense_id = $1
	`, input.ExpenseID).Scan(&billFile)

	if err == sql.ErrNoRows {
		http.Error(w, `{"success": false, "message": "No expense found with the given expense_id"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"success": false, "message": "Database query error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	result, err := dbconfig.Database.Exec(`
		DELETE FROM manage_expenditure WHERE expense_id = $1
	`, input.ExpenseID)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error deleting expense: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No expense found to delete"}`, http.StatusNotFound)
		return
	}

	if billFile.Valid && billFile.String != "" {
		wd, _ := os.Getwd()

		filePath := filepath.Join(wd, "expense_bills", billFile.String)

		if _, err := os.Stat(filePath); err == nil {
			if rmErr := os.Remove(filePath); rmErr == nil {
				fmt.Print("Expense bill deleted")
			}
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":      true,
		"message":      "Expense deleted successfully",
		"expense_id":   input.ExpenseID,
		"deleted_at":   time.Now(),
		"bill_deleted": billFile.Valid && billFile.String != "",
	})
}
