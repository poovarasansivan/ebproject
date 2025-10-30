package models

type UserModel struct {
	UserID         int    `json:"user_id"`
	Username       string `json:"user_name"`
	Email          string `json:"user_email"`
	PhoneNumber    string `json:"phone_no"`
	Role           string `json:"role"`
	AccountStatus  string `json:"account_status"`
}