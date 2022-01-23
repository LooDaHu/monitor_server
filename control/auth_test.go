package control

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"monitor_server/message"
	"monitor_server/model"
	"monitor_server/utils"
	"net/http"
	"net/http/httptest"
)

func TestControllerAuth_AppCodeCheckRequired(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := model.NewMockUserer(ctrl)

	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_Success", t, func() {
		req := &message.LoginReq{
			Username: "hello",
			Password: "world",
		}
		info, err := m.RetrieveUserInfo(bson.M{})
		res := CheckUserInfo(req, info)
		convey.So(res, convey.ShouldEqual, true)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestControllerAuth_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := model.NewMockUserer(ctrl)
	// 参数校验失败
	convey.Convey("TestUser_CreateUserInfo_参数校验失败", t, func() {
		var err utils.Error
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
		}
		MockJsonPost(ctx, map[string]interface{}{"foo": "bar"})
		LoginTest(ctx, m)
		json.Unmarshal(w.Body.Bytes(), &err)
		fmt.Println(string(w.Body.Bytes()))
		convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(err.Code, convey.ShouldEqual, utils.ParamsCheckErrorCode)
	})
	// 认证失败
	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world123",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_认证失败", t, func() {
		var err utils.Error
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
		}
		MockJsonPost(ctx, map[string]interface{}{
			"username": "hello",
			"password": "world1",
			"type":     "admin"})
		LoginTest(ctx, m)
		json.Unmarshal(w.Body.Bytes(), &err)
		fmt.Println(string(w.Body.Bytes()))
		convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(err.Code, convey.ShouldEqual, utils.AuthFailErrorCode)
	})
	// 认证成功
	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world123",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_认证成功", t, func() {
		var err utils.Error
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
		}
		MockJsonPost(ctx, map[string]interface{}{
			"username": "hello",
			"password": "world123",
			"type":     "admin"})
		LoginTest(ctx, m)
		json.Unmarshal(w.Body.Bytes(), &err)
		fmt.Println(string(w.Body.Bytes()))
		convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(err.Code, convey.ShouldEqual, utils.SuccessCode)
	})
}

func TestControllerAuth_LoginRequired(t *testing.T) {
	// 涉及session与中间件，跳过单元测试
}

func TestControllerAuth_Logout(t *testing.T) {
	// 涉及session与中间件，跳过单元测试
}

func TestCheckUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := model.NewMockUserer(ctrl)

	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_成功", t, func() {
		req := &message.LoginReq{
			Username: "hello",
			Password: "world",
		}
		info, err := m.RetrieveUserInfo(bson.M{})
		res := CheckUserInfo(req, info)
		convey.So(res, convey.ShouldEqual, true)
		convey.So(err, convey.ShouldEqual, nil)
	})
	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_密码不匹配", t, func() {
		req := &message.LoginReq{
			Username: "hello",
			Password: "world1",
		}
		info, _ := m.RetrieveUserInfo(bson.M{})
		res := CheckUserInfo(req, info)
		convey.So(res, convey.ShouldEqual, false)
	})
	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_账户不匹配", t, func() {
		req := &message.LoginReq{
			Username: "hello1",
			Password: "world",
		}
		info, _ := m.RetrieveUserInfo(bson.M{})
		res := CheckUserInfo(req, info)
		convey.So(res, convey.ShouldEqual, false)
	})
	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&model.User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "world",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_CreateUserInfo_密码账户都不匹配", t, func() {
		req := &message.LoginReq{
			Username: "hello1",
			Password: "world1",
		}
		info, _ := m.RetrieveUserInfo(bson.M{})
		res := CheckUserInfo(req, info)
		convey.So(res, convey.ShouldEqual, false)
	})
}

//Login 用户登录 单元测试用
func LoginTest(c *gin.Context, m *model.MockUserer) {
	logInfo := new(message.LoginReq)
	//session := sessions.Default(c)

	//校验参数
	if err := c.ShouldBindWith(&logInfo, binding.JSON); err != nil {
		c.JSON(http.StatusOK, utils.ParamsCheckError)
		return
	}
	//utils.SugarLogger.Info("Login Call", logInfo)
	filter := bson.M{
		"username": logInfo.Username,
		"password": logInfo.Password,
	}
	user, err := m.RetrieveUserInfo(filter) // 从数据库检查用户信息
	if err != nil {
		c.JSON(http.StatusOK, utils.SystemError)
		return
	}
	isPass := CheckUserInfo(logInfo, user)
	if !isPass {
		c.JSON(http.StatusOK, utils.AuthFailError)
		return
	}
	//session.Set(UserKey, user.ID) // 保存用户session
	//session.Set(TypeKey, user.Type)
	//if err := session.Save(); err != nil {
	//	c.JSON(http.StatusInternalServerError, utils.SystemError)
	//	return
	//}
	c.JSON(http.StatusOK, utils.Success)
}

// 构造Mock Json Post请求
func MockJsonPost(c *gin.Context /* the test context */, content interface{}) {
	c.Request.Method = "POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}
