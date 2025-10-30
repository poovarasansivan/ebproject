package manageusers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/dbconfig"
)

func UpdateUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		UserID int    `json:"user_id"`
		Role   string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.UserID <= 0 || input.Role == "" {
		http.Error(w, "Invalid or missing user_id or role", http.StatusBadRequest)
		return
	}

	if dbconfig.Database == nil {
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}

	result, err := dbconfig.Database.Exec(`
		UPDATE users
		SET role = $1
		WHERE user_id = $2
	`, input.Role, input.UserID)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No user found with the given user_id", http.StatusNotFound)
			return
		}
		http.Error(w, "Database update error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Error fetching update result: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No user found with the given user_id", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "User role updated successfully",
	})
}
