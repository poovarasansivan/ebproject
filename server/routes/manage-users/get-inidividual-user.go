package manageusers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetIndividualUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		UserID int `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.UserID <= 0 {
		http.Error(w, "Invalid or missing user_id", http.StatusBadRequest)
		return
	}

	var user models.UserModel
	err := dbconfig.Database.QueryRow(`
		SELECT user_id, user_name, user_email, phone_no, role, acc_status
		FROM users
		WHERE user_id = $1
	`, input.UserID).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.PhoneNumber,
		&user.Role,
		&user.AccountStatus,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    user,
	})
}
