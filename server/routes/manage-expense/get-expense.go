package manageexpense

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
)

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use GET."}`, http.StatusMethodNotAllowed)
		return
	}

	query := `
		SELECT 
			me.expense_id,
			me.labour_id,
			labour.user_name AS labour_name,
			me.manager_id,
			manager.user_name AS manager_name,
			me.expense_description,
			me.total_expense,
			me.bill_docs,
			me.created_at,
			me.updated_at
		FROM manage_expenditure me
		LEFT JOIN users labour ON CAST(me.labour_id AS INT) = labour.user_id
		LEFT JOIN users manager ON CAST(me.manager_id AS INT) = manager.user_id
		ORDER BY me.created_at DESC;
	`

	rows, err := dbconfig.Database.Query(query)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database query error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type ExpenseWithNames struct {
		ExpenseID          int     `json:"expense_id"`
		LabourID           int     `json:"labour_id"`
		LabourName         string  `json:"labour_name"`
		ManagerID          string  `json:"manager_id"`
		ManagerName        string  `json:"manager_name"`
		ExpenseDescription string  `json:"expense_description"`
		TotalExpense       string  `json:"total_expense"`
		BillDocs           string  `json:"bill_docs"`
		CreatedAt          string  `json:"created_at"`
		UpdatedAt          string  `json:"updated_at"`
	}

	var expenses []ExpenseWithNames

	for rows.Next() {
		var e ExpenseWithNames
		err := rows.Scan(
			&e.ExpenseID,
			&e.LabourID,
			&e.LabourName,
			&e.ManagerID,
			&e.ManagerName,
			&e.ExpenseDescription,
			&e.TotalExpense,
			&e.BillDocs,
			&e.CreatedAt,
			&e.UpdatedAt,
		)
		if err != nil {
			http.Error(w, `{"success": false, "message": "Error scanning expense data: `+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
		expenses = append(expenses, e)
	}

	if len(expenses) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": "No expenses found",
			"data":    []ExpenseWithNames{},
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"count":   len(expenses),
		"data":    expenses,
	})
}
