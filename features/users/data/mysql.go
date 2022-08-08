package data

import (
	"newsapi/features/users"
	"newsapi/plugins"

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

func (repo *mysqlUserRepository) LoginUser(email, password string) (idUser int) {
	var data User
	result := repo.db.First(&data, "email = ?", email)
	if result.Error != nil {
		return 0
	}
	if !plugins.CheckPasswordHash(password, data.Password) {
		return 0
	}
	return int(data.ID)
}

func (repo *mysqlUserRepository) SelectProfile(idUser int) (users.Core, error) {
	var data User
	result := repo.db.First(&data, "id = ?", idUser)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return data.toCore(), nil
}

func (repo *mysqlUserRepository) InsertUser(input users.Core) (int, error) {
	var newPass, _ = plugins.HashPassword(input.Password)
	input.Password = newPass
	var data = fromCore(input)
	result := repo.db.Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
