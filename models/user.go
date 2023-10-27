package models

// User represents a user entity
type SignupUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
