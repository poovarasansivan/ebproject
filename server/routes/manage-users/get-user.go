package manageusers

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := dbconfig.Database.Query(`
		SELECT user_id, user_name, user_email, phone_no, role, acc_status
		FROM users
		ORDER BY user_id ASC
	`)
	if err != nil {
		http.Error(w, "Database query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.UserModel

	for rows.Next() {
		var u models.UserModel
		err := rows.Scan(
			&u.UserID,
			&u.Username,
			&u.Email,
			&u.PhoneNumber,
			&u.Role,
			&u.AccountStatus,
		)
		if err != nil {
			http.Error(w, "Error scanning user data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"count":   len(users),
		"data":    users,
	})
}
