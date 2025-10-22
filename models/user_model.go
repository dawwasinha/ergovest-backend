package models

// Struct untuk login user (admin, operator, supervisor)
type User struct {
    Username string `json:"username"`
    Role     string `json:"role"`
}