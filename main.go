package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"monitor_server/services"

	. "monitor_server/utils"
)

type allService struct {
	Auth services.AuthSrv
	Info services.SystemInfoSrv
}

func main() {
	InitLogger()
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		SugarLogger.Fatal("Unable to start", err)
	}
}

func engine() *gin.Engine {
	srv := new(allService)
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.POST("/login", srv.Auth.Ctrl.Login)  // 登入
	r.GET("/logout", srv.Auth.Ctrl.Logout) // 登出

	rInfo := r.Group("/sysinfo")
	rInfo.Use(srv.Auth.Ctrl.LoginRequired)
	{
		rInfo.POST("/retrieve", srv.Info.Ctrl.QuerySystemInfo) // 查询上报信息

	}

	rAgent := r.Group("/agent")
	rAgent.Use(srv.Auth.Ctrl.AppCodeCheckRequired)
	{
		rAgent.POST("/upload", srv.Info.Ctrl.UploadSystemInfo) // agent上报信息
	}

	return r
}
