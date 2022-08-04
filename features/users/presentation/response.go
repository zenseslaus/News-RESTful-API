package presentation

import "newsapi/features/users"

type ResponseUser struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func FromCore(data users.Core) ResponseUser {
	return ResponseUser{
		ID:       data.ID,
		FullName: data.FullName,
		UserName: data.UserName,
		Email:    data.Email,
	}
}
