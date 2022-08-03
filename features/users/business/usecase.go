package business

import "newsapi/features/users"

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(data users.Data) users.Business {
	return &userUsecase{
		userData: data,
	}
}

func (uc *userUsecase) GetProfile(idUser int) (users.Core, error) {
	result, err := uc.userData.SelectProfile(idUser)
	return result, err
}
