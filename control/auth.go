package control

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"

	"net/http"
	"strings"

	"monitor_server/message"
	"monitor_server/model"
	. "monitor_server/utils"
)

const (
	UserKey = "user"
	TypeKey = "type"
)

type ControllerAuth struct {
	user model.User
}

// Login 用户登录
func (ca ControllerAuth) Login(c *gin.Context) {
	logInfo := new(message.LoginReq)
	session := sessions.Default(c)
	SugarLogger.Info("Login Call", logInfo)

	//校验参数
	if err := c.ShouldBindWith(&logInfo, binding.JSON); err != nil {
		c.JSON(http.StatusOK, ParamsCheckError)
	}
	filter := bson.M{
		"username": logInfo.Username,
		"password": logInfo.Password,
	}
	user, err := ca.user.RetrieveUserInfo(filter) // 从数据库检查用户信息
	if err != nil {
		c.JSON(http.StatusOK, SystemError)
		return
	}
	isPass := CheckUserInfo(logInfo, user)
	if !isPass {
		c.JSON(http.StatusOK, AuthFailError)
		return
	}
	session.Set(UserKey, user.ID) // 保存用户session
	session.Set(TypeKey, user.Type)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, SystemError)
		return
	}
	c.JSON(http.StatusOK, Success)
}

// Logout 用户登出
func (ca ControllerAuth) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	if user == nil {
		c.JSON(http.StatusBadRequest, SystemError)
		return
	}
	session.Delete(UserKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, SystemError)
		return
	}
	c.JSON(http.StatusOK, Success)
}

// LoginRequired 检查session中间件
func (ca ControllerAuth) LoginRequired(c *gin.Context) {
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

// AppCodeCheckRequired 检查App Code
func (ca ControllerAuth) AppCodeCheckRequired(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	keyValue := strings.Split(auth, " ")
	if len(keyValue) != 2 {
		// code校验不合法 中断请求
		c.AbortWithStatusJSON(http.StatusOK, AuthFailError)
		return
	}

	appCode, err := model.RetrieveAppCode(bson.M{
		"app_id":   keyValue[0],
		"app_code": keyValue[1],
	})
	if appCode == nil {
		c.AbortWithStatusJSON(http.StatusOK, AuthFailError)
		return
	}
	if err != nil {
		SugarLogger.Error("mongoDB error@ AppCodeCheckRequired", err)
		c.AbortWithStatusJSON(http.StatusOK, SystemError)
		return
	}
	c.Next()
}

func CheckUserInfo(req *message.LoginReq, user *model.User) bool {
	return req.Username == user.Username && req.Password == user.Password
}
