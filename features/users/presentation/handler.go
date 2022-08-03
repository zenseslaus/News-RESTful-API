package presentation

import (
	"net/http"
	"newsapi/features/users"
	"strconv"

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

func (hand *UserHandler) GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("idUser")
		idInt, errId := strconv.Atoi(id)
		if errId != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error bind",
			})
			return
		}
		result, err := hand.userBusiness.GetProfile(idInt)
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
