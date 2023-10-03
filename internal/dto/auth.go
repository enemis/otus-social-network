package dto

import "time"

type SignUpInput struct {
	Name      string    `json:"name" binding:"required" required:"$field is required"`
	Surname   string    `json:"surname" binding:"required" required:"$field is required"`
	Birthday  time.Time `json:"birthday" binding:"required" time_format:"2006-01-02"`
	Email     string    `json:"email" binding:"required,email" required:"$field is required"`
	Biography string    `json:"biography" binding:"omitempty"`
	City      string    `json:"city" binding:"required" required:"$field is required"`
	Password  string    `json:"password" binding:"required" required:"$field is required"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


