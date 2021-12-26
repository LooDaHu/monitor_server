package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"encoding/json"
	"net/http"

	"monitor_server/message"
	"monitor_server/model"
	. "monitor_server/utils"
)

func UploadSystemInfo(c *gin.Context) {
	SugarLogger.Info("UploadSystemInfo Call")
	req := new(message.UploadSystemInfoReq)
	//校验参数
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		SugarLogger.Error("Params Check Failed @UploadSystemInfo", err)
		c.JSON(http.StatusOK, ParamsCheckError)
	}
	orders, err := model.RetrieveOrder(info.User)
	if err != nil {
		SugarLogger.Error("RetrieveOrder Failed @UploadSystemInfo", err)
		c.JSON(http.StatusOK, SystemError)
		return
	}
	data, _ := json.Marshal(orders)
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"data":    data,
	})
}
