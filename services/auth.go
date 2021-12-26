package services

import (
	"github.com/gin-gonic/gin"
	"monitor_server/control"
)

type AuthService interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	LoginRequired(c *gin.Context)
	AppCodeCheckRequired(c *gin.Context)
}

type AuthSrv struct {
	Ctrl control.ControllerAuth
}
