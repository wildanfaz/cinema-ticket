package models

import "time"

type (
	User struct {
		Id        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Role      string    `json:"role"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Login struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
)

func (m *User) ToLocal() {
	m.CreatedAt = m.CreatedAt.Local()
	m.UpdatedAt = m.UpdatedAt.Local()
}
