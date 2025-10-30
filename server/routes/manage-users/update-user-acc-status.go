package manageusers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/dbconfig"
)

type UpdateUserStatusRequest struct {
	UserID        int    `json:"user_id"`
	AccountStatus string `json:"acc_status"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func UpdateUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid request method",
		})
		return
	}

	var input UpdateUserStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid JSON format",
			Error:   err.Error(),
		})
		return
	}
	defer r.Body.Close()

	if input.UserID <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid user_id",
		})
		return
	}
	if input.AccountStatus != "Active" && input.AccountStatus != "Inactive" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid account_status (must be 'Active' or 'Inactive')",
		})
		return
	}

	if dbconfig.Database == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Database not initialized",
		})
		return
	}

	result, err := dbconfig.Database.Exec(`
		UPDATE users
		SET acc_status = $1, updated_at = NOW()
		WHERE user_id = $2
	`, input.AccountStatus, input.UserID)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIResponse{
				Success: false,
				Message: "No user found with the given user_id",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Database update error",
			Error:   err.Error(),
		})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Error fetching update result",
			Error:   err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "No user found with the given user_id",
		})
		return
	}

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "User account status updated successfully",
		Data: map[string]interface{}{
			"user_id":        input.UserID,
			"account_status": input.AccountStatus,
		},
	})
}
