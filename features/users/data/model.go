package data

import (
	"newsapi/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	UserName string `gorm:"unique" json:"user_name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        data.ID,
		FullName:  data.FullName,
		UserName:  data.UserName,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func fromCore(data users.Core) User {
	return User{
		FullName: data.FullName,
		UserName: data.UserName,
		Email:    data.Email,
		Password: data.Password,
	}
}
