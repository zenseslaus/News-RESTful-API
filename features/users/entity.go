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
}

type Business interface {
	GetProfile(idUser int) (data Core, err error)
}

type Data interface {
	SelectProfile(idUser int) (data Core, err error)
}
