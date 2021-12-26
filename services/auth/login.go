package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"

	"net/http"

	"monitor_server/message"
	"monitor_server/model"
	. "monitor_server/utils"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	UserKey = "user"
	TypeKey = "type"
)

// LoginRequired 检查session中间件
func LoginRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	if user == nil {
		// 未登录，阻断请求
		c.AbortWithStatusJSON(http.StatusOK, AuthFailError)
		return
	}
	// 继续
	c.Next()
}

// AppCodeCheckRequired 检查App Coe
func AppCodeCheckRequired(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	keyValue := strings.Split(auth, " ")
	if len(keyValue) != 2 {
		// 未登录，阻断请求
		c.AbortWithStatusJSON(http.StatusOK, AuthFailError)
		return
	}

}

// Login 用户登录
func Login(c *gin.Context) {
	logInfo := new(message.LoginReq)
	session := sessions.Default(c)
	SugarLogger.Info("Login Call", logInfo)

	//校验参数
	if err := c.ShouldBindWith(&logInfo, binding.JSON); err != nil {
		c.JSON(http.StatusOK, ParamsCheckError)
	}
	isPass, id, userType, err := model.CheckUserInfo(logInfo.Username,
		logInfo.Password) // 从数据库检查用户信息
	if err != nil {
		c.JSON(http.StatusOK, SystemError)
		return
	}
	if !isPass {
		c.JSON(http.StatusOK, AuthFailError)
		return
	}

	session.Set(UserKey, id) // 保存用户session
	session.Set(TypeKey, userType)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, SystemError)
		return
	}
	c.JSON(http.StatusOK, Success)
}
