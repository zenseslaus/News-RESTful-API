package data

import (
	"newsapi/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserData(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) SelectProfile(idUser int) (users.Core, error) {
	var data User
	result := repo.db.First(&data, "id = ?", idUser)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return data.toCore(), nil
}
