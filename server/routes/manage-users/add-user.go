package manageusers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dbconfig"
	"server/models"
)

type AddUserResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.UserModel
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if newUser.Username == "" || newUser.Email == "" || newUser.Role == "" {
		http.Error(w, "Missing required fields (username, email, role)", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO users (user_name, user_email, phone_no, role, acc_status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id, user_name, user_email, phone_no, role, acc_status;
	`

	err := dbconfig.Database.QueryRow(
		query,
		newUser.Username,
		newUser.Email,
		newUser.PhoneNumber,
		newUser.Role,
		newUser.AccountStatus,
	).Scan(
		&newUser.UserID,
		&newUser.Username,
		&newUser.Email,
		&newUser.PhoneNumber,
		&newUser.Role,
		&newUser.AccountStatus,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %v", err), http.StatusInternalServerError)
		return
	}

	response := AddUserResponse{
		Success: true,
		Message: "User added successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
