package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"monitor_server/services/auth"
	"monitor_server/services/system"
	. "monitor_server/utils"
)

func main() {
	InitLogger()
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		SugarLogger.Fatal("Unable to start", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.POST("/login", auth.Login)  // 登入
	r.GET("/logout", auth.Logout) // 登出

	rInfo := r.Group("/sysinfo")
	rInfo.Use(auth.LoginRequired)
	{
		rInfo.POST("/retrieve", system.UploadSystemInfo) // 查询上报信息

	}

	rAgent := r.Group("/agent")
	rAgent.Use(auth.AppCodeCheckRequired)
	{
		rAgent.POST("/upload", system.UploadSystemInfo) // agent上报信息
	}

	return r
}
