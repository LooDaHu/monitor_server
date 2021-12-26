package control

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"monitor_server/message"
	"monitor_server/model"
	. "monitor_server/utils"
)

type ControllerSystemInfo struct {
	sysInfo model.SysInfo
}

func (c2 ControllerSystemInfo) UploadSystemInfo(c *gin.Context) {
	SugarLogger.Info("UploadSystemInfo Call")
	req := new(message.UploadSystemInfoReq)
	//校验参数
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		SugarLogger.Error("Params Check Failed @UploadSystemInfo", err)
		c.JSON(http.StatusOK, ParamsCheckError)
	}
	info, err := InfoBuilder(req)
	if err != nil {
		SugarLogger.Error(" Builder Info Failed @UploadSystemInfo", err)
		c.JSON(http.StatusOK, SystemError)
		return
	}
	err = c2.sysInfo.CreateSystemInfo(info)
	if err != nil {
		SugarLogger.Error("RetrieveOrder Failed @UploadSystemInfo", err)
		c.JSON(http.StatusOK, SystemError)
		return
	}
	c.JSON(http.StatusOK, Success)
}

func (c2 ControllerSystemInfo) QuerySystemInfo(c *gin.Context) {
	SugarLogger.Info("QuerySystemInfo Call")
	req := new(message.QuerySystemInfoReq)
	//校验参数
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		SugarLogger.Error("Params Check Failed @QuerySystemInfo", err)
		c.JSON(http.StatusOK, ParamsCheckError)
	}
	filter := FilterBuilder(req)
	systemInfo, err := c2.sysInfo.RetrieveSystemInfo(filter)
	if err != nil || systemInfo == nil {
		SugarLogger.Error("MongoDB Ops Failed @QuerySystemInfo", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "系统错误"})
		return
	}
	data, _ := json.Marshal(systemInfo)
	c.JSON(http.StatusOK, gin.H{
		"code":    Success.Code,
		"message": Success.Message,
		"data":    data,
	})
}

func InfoBuilder(req *message.UploadSystemInfoReq) (*model.SysInfo, error) {
	var sysInfo model.SystemInfo
	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &sysInfo)
	if err != nil {
		return nil, err
	}
	return &model.SysInfo{
		SystemInfo: sysInfo,
		CreateTime: primitive.Timestamp{
			T: uint32(time.Now().Unix()),
			I: 0,
		},
	}, nil
}

func FilterBuilder(req *message.QuerySystemInfoReq) bson.M {
	return bson.M{
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
}
