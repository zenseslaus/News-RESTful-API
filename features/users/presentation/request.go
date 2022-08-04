package presentation

import "newsapi/features/users"

type InsertUser struct {
	FullName string `json:"full_name" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (data *InsertUser) ToCore() users.Core {
	return users.Core{
		FullName: data.FullName,
		UserName: data.UserName,
		Email:    data.Email,
		Password: data.Password,
	}
}
