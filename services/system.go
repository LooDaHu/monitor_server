package services

import (
	"github.com/gin-gonic/gin"
	"monitor_server/control"
)

type SystemService interface {
	UploadSystemInfo(c *gin.Context)
	QuerySystemInfo(c *gin.Context)
}

type SystemInfoSrv struct {
	Ctrl control.ControllerSystemInfo
}
