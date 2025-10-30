package manageusers

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"strconv"
	"strings"
)

func UpdateUserDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		UserID      int     `json:"user_id"`
		Username    *string `json:"user_name,omitempty"`
		Email       *string `json:"user_email,omitempty"`
		PhoneNumber *string `json:"phone_no,omitempty"`
		Role        *string `json:"role,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.UserID <= 0 {
		http.Error(w, "Missing or invalid user_id", http.StatusBadRequest)
		return
	}

	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if input.Username != nil && *input.Username != "" {
		setClauses = append(setClauses, "user_name = $"+strconv.Itoa(argPos))
		args = append(args, *input.Username)
		argPos++
	}

	if input.Email != nil && *input.Email != "" {
		setClauses = append(setClauses, "user_email = $"+strconv.Itoa(argPos))
		args = append(args, *input.Email)
		argPos++
	}

	if input.PhoneNumber != nil && *input.PhoneNumber != "" {
		setClauses = append(setClauses, "phone_no = $"+strconv.Itoa(argPos))
		args = append(args, *input.PhoneNumber)
		argPos++
	}

	if input.Role != nil && *input.Role != "" {
		setClauses = append(setClauses, "role = $"+strconv.Itoa(argPos))
		args = append(args, *input.Role)
		argPos++
	}

	if len(setClauses) == 0 {
		http.Error(w, "No fields provided to update", http.StatusBadRequest)
		return
	}

	setClauses = append(setClauses, "updated_at = NOW()")

	query := `
		UPDATE users
		SET ` + strings.Join(setClauses, ", ") + `
		WHERE user_id = $` + strconv.Itoa(argPos)
	args = append(args, input.UserID)

	result, err := dbconfig.Database.Exec(query, args...)
	if err != nil {
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
		"message": "User details updated successfully",
	})
}
