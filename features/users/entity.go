package users

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Core struct {
	ID        uint
	FullName  string
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Handler interface {
	GetUserProfile() gin.HandlerFunc
	PostNewUser() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type Business interface {
	GetProfile(idUser int) (data Core, err error)
	PostUser(data Core) (row int, err error)
	LoginUser(email, password string) (token string)
}

type Data interface {
	SelectProfile(idUser int) (data Core, err error)
	InsertUser(data Core) (row int, err error)
	LoginUser(email, password string) (token string)
}
