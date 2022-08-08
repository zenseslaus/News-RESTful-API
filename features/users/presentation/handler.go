package presentation

import (
	"net/http"
	"newsapi/features/users"
	"newsapi/helpers"
	"newsapi/middlewares"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) users.Handler {
	return &UserHandler{
		userBusiness: business,
	}
}
func (hand *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login LoginUser
		errBind := c.Bind(&login)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseFailed("error bind"))
			return
		}
		token := hand.userBusiness.LoginUser(login.Email, login.Password)
		c.JSON(http.StatusOK, helpers.ResponseSuccesWithData("success", token))
	}
}

func (hand *UserHandler) GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		idUser, errId := middlewares.ExtractToken(*c)
		if errId != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed token",
			})
			return
		}
		result, err := hand.userBusiness.GetProfile(idUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get data",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    FromCore(result),
		})
	}
}

func (hand *UserHandler) PostNewUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataReq InsertUser
		errBind := c.Bind(&dataReq)
		if errBind != nil {
			c.JSON(http.StatusBadRequest, helpers.ResponseFailed("error bind"))
			return
		}
		row, err := hand.userBusiness.PostUser(dataReq.ToCore())
		if row == 0 {
			c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert data"))
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, helpers.ResponseFailed("failed to insert data"))
			return
		}
		c.JSON(http.StatusOK, helpers.ResponseSuccesNoData("success"))
	}
}
