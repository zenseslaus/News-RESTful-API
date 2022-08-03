package factory

import (
	_userBusiness "newsapi/features/users/business"
	_userData "newsapi/features/users/data"
	_userPresentation "newsapi/features/users/presentation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitFactory(r *gin.Engine, dbConn *gorm.DB) {
	UserData := _userData.NewUserData(dbConn)
	UserBusiness := _userBusiness.NewUserBusiness(UserData)
	UserPresentation := _userPresentation.NewUserHandler(UserBusiness)
	_userPresentation.RouteUser(r, UserPresentation)
}
