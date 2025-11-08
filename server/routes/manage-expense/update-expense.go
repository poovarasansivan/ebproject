package manageexpense

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dbconfig"
	"strings"
	"time"
)

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST or PUT."}`, http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		ExpenseID          int     `json:"expense_id"`
		LabourID           *int    `json:"labour_id,omitempty"`
		ManagerID          *string `json:"manager_id,omitempty"`
		ExpenseDescription *string `json:"expense_description,omitempty"`
		TotalExpense       *string `json:"total_expense,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.ExpenseID <= 0 {
		http.Error(w, `{"success": false, "message": "expense_id is required and must be valid"}`, http.StatusBadRequest)
		return
	}

	setClauses := []string{}
	params := []interface{}{}
	paramIndex := 1

	if input.LabourID != nil {
		setClauses = append(setClauses, fmt.Sprintf("labour_id = $%d", paramIndex))
		params = append(params, *input.LabourID)
		paramIndex++
	}
	if input.ManagerID != nil {
		setClauses = append(setClauses, fmt.Sprintf("manager_id = $%d", paramIndex))
		params = append(params, *input.ManagerID)
		paramIndex++
	}
	if input.ExpenseDescription != nil {
		setClauses = append(setClauses, fmt.Sprintf("expense_description = $%d", paramIndex))
		params = append(params, *input.ExpenseDescription)
		paramIndex++
	}
	if input.TotalExpense != nil {
		setClauses = append(setClauses, fmt.Sprintf("total_expense = $%d", paramIndex))
		params = append(params, *input.TotalExpense)
		paramIndex++
	}

	if len(setClauses) == 0 {
		http.Error(w, `{"success": false, "message": "No fields provided to update"}`, http.StatusBadRequest)
		return
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", paramIndex))
	params = append(params, time.Now())
	paramIndex++

	query := fmt.Sprintf(`UPDATE manage_expenditure SET %s WHERE expense_id = $%d`,
		strings.Join(setClauses, ", "), paramIndex)

	params = append(params, input.ExpenseID)

	result, err := dbconfig.Database.Exec(query, params...)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database update error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error fetching update result"}`, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No expense found with the given expense_id"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    true,
		"message":    "Expense updated successfully",
		"expense_id": input.ExpenseID,
		"updated_at": time.Now(),
	})
}
