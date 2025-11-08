package manageexpense

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/handlebills"
	"time"
)

func AddExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST."}`, http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, `{"success": false, "message": "Error parsing form data: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	labourID := r.FormValue("labour_id")
	managerID := r.FormValue("manager_id")
	expenseDescription := r.FormValue("expense_description")
	totalExpense := r.FormValue("total_expense")

	if labourID == "" || managerID == "" || totalExpense == "" {
		http.Error(w, `{"success": false, "message": "labour_id, manager_id, and total_expense are required"}`, http.StatusBadRequest)
		return
	}

	fileName, err := handlebills.UploadExpenseFile(r, "bill_docs")
	if err != nil {
		http.Error(w, `{"success": false, "message": "File upload failed: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	query := `
		INSERT INTO manage_expenditure 
		(labour_id, manager_id, expense_description, total_expense, bill_docs, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING expense_id;
	`

	var expenseID int
	err = dbconfig.Database.QueryRow(
		query,
		labourID,
		managerID,
		expenseDescription,
		totalExpense,
		fileName,
	).Scan(&expenseID)

	if err != nil {
		http.Error(w, `{"success": false, "message": "Database insert error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    true,
		"message":    "Expense added successfully",
		"expense_id": expenseID,
		"bill_docs":  fileName,
		"created_at": time.Now(),
	})
}
