package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("VJzY3r8Zq9sXc2v5y8zD1fG4hJkL0mN2p")

type SigninRequest struct {
	UserEmail string `json:"user_email"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Invalid request method",
		})
		return
	}

	var input SigninRequest
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

	if input.UserEmail == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Email is required",
		})
		return
	}

	var dbUser models.UserModel
	err := dbconfig.Database.QueryRow(`
		SELECT user_id, user_name, user_email, phone_no, role, acc_status
		FROM users
		WHERE user_email = $1
	`, input.UserEmail).Scan(
		&dbUser.UserID,
		&dbUser.Username,
		&dbUser.Email,
		&dbUser.PhoneNumber,
		&dbUser.Role,
		&dbUser.AccountStatus,
	)

	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "User not found",
		})
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Database error",
			Error:   err.Error(),
		})
		return
	}

	if dbUser.AccountStatus != "Active" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "User account is inactive. Please contact administrator.",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dbUser.UserID,
		"email":   dbUser.Email,
		"role":    dbUser.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(APIResponse{
			Success: false,
			Message: "Failed to generate authentication token",
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(APIResponse{
		Success: true,
		Message: "Signin successful",
		Token:   tokenString,
	})
}
