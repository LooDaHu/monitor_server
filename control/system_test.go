package control

import (
	"encoding/json"
	"fmt"
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
	"testing"
)

func TestControllerSystemInfo_QuerySystemInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := model.NewMockSysInfoer(ctrl)
	// 参数校验失败
	convey.Convey("TestControllerSystemInfo_QuerySystemInfo_参数校验失败", t, func() {
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
		QuerySystemInfoTest(ctx, m)
		json.Unmarshal(w.Body.Bytes(), &err)
		fmt.Println(string(w.Body.Bytes()))
		convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(err.Code, convey.ShouldEqual, utils.ParamsCheckErrorCode)
	})
	m.EXPECT().RetrieveSystemInfo(gomock.Any()).Return(&model.SysInfo{
		SystemInfo: model.SystemInfo{},
		CreateTime: primitive.Timestamp{},
	}, nil)
	convey.Convey("TestControllerSystemInfo_QuerySystemInfo_参数校验失败", t, func() {
		var err utils.Error
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
		}
		MockJsonPost(ctx, map[string]interface{}{
			"host_name":  "test",
			"os":         "Windows",
			"start_time": 1600000000,
			"end_time":   1600000001,
		})
		QuerySystemInfoTest(ctx, m)
		json.Unmarshal(w.Body.Bytes(), &err)
		fmt.Println(string(w.Body.Bytes()))
		convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
		convey.So(err.Code, convey.ShouldEqual, utils.SuccessCode)
	})
}

func TestControllerSystemInfo_UploadSystemInfo(t *testing.T) {
	convey.Convey("TestInfoBuilder_构建参数成功", t, func() {
		req := &message.UploadSystemInfoReq{
			Host:    message.HostInfo{},
			Network: message.NetworkInfo{},
			CPU:     message.CPUInfo{},
			Memory:  message.MemoryInfo{},
			Disk:    message.DiskInfo{},
		}
		info, err := InfoBuilder(req)
		convey.So(info, convey.ShouldNotHaveSameTypeAs, &model.SystemInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestFilterBuilder(t *testing.T) {
	convey.Convey("TestFilterBuilder_构建过滤器成功", t, func() {
		req := &message.QuerySystemInfoReq{
			HostName:  "hello",
			OS:        "Windows",
			StartTime: 1610000000,
			EndTime:   1610000001,
		}
		filter := FilterBuilder(req)
		filter1 := bson.M{
			"host_name": req.HostName,
			"os":        req.OS,
			"create_time": bson.M{
				"$gte": primitive.Timestamp{
					T: req.StartTime,
					I: 0,
				},
				"$lte": primitive.Timestamp{
					T: req.EndTime,
					I: 0,
				},
			},
		}
		convey.So(filter, convey.ShouldResemble, filter1)
	})
}

func TestInfoBuilder(t *testing.T) {
	convey.Convey("TestInfoBuilder_构建参数成功", t, func() {
		req := &message.UploadSystemInfoReq{
			Host:    message.HostInfo{},
			Network: message.NetworkInfo{},
			CPU:     message.CPUInfo{},
			Memory:  message.MemoryInfo{},
			Disk:    message.DiskInfo{},
		}
		info, err := InfoBuilder(req)
		convey.So(info, convey.ShouldNotHaveSameTypeAs, &model.SystemInfo{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func QuerySystemInfoTest(c *gin.Context, m *model.MockSysInfoer) {
	//utils.SugarLogger.Info("QuerySystemInfo Call")
	req := new(message.QuerySystemInfoReq)
	//校验参数
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		//utils.SugarLogger.Error("Params Check Failed @QuerySystemInfo", err)
		c.JSON(http.StatusOK, utils.ParamsCheckError)
		return
	}
	filter := FilterBuilder(req)
	systemInfo, err := m.RetrieveSystemInfo(filter)
	if err != nil || systemInfo == nil {
		//utils.SugarLogger.Error("MongoDB Ops Failed @QuerySystemInfo", err)
		c.JSON(http.StatusInternalServerError, utils.SystemError)
		return
	}
	data, _ := json.Marshal(systemInfo)
	c.JSON(http.StatusOK, gin.H{
		"code":    utils.Success.Code,
		"message": utils.Success.Message,
		"data":    data,
	})
}
