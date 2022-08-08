package business

import (
	"newsapi/features/users"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(data users.Data) users.Business {
	return &userUsecase{
		userData: data,
	}
}

func (uc *userUsecase) LoginUser(email, password string) int {
	return uc.userData.LoginUser(email, password)
}

func (uc *userUsecase) GetProfile(idUser int) (users.Core, error) {
	result, err := uc.userData.SelectProfile(idUser)
	return result, err
}

func (uc *userUsecase) PostUser(input users.Core) (int, error) {
	row, err := uc.userData.InsertUser(input)
	return row, err
}
