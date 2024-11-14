package models

type User struct {
	User_id  string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"-"` // Exclude from response JSON, we'll hash it for storage
	Role     string `json:"role"`
}
