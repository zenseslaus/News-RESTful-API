package presentation

import (
	"newsapi/features/users"

	"github.com/gin-gonic/gin"
)

func RouteUser(r *gin.Engine, uc users.Handler) {
	r.GET("/users", uc.GetUserProfile())
	r.POST("/users", uc.PostNewUser())
	r.POST("/login", uc.Login())
}
