package model

import (
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestSysInfo_CreateSystemInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockSysInfoer(ctrl)

	m.EXPECT().CreateSystemInfo(gomock.Any()).Return(nil)
	convey.Convey("TestSysInfo_CreateSystemInfo", t, func() {
		err := m.CreateSystemInfo(&SysInfo{
			SystemInfo: SystemInfo{},
			CreateTime: primitive.Timestamp{T: uint32(time.Now().Unix())},
		})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestSysInfo_DeleteSystemInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockSysInfoer(ctrl)

	m.EXPECT().DeleteSystemInfo(gomock.Any()).Return(nil)
	convey.Convey("TestSysInfo_DeleteSystemInfo", t, func() {
		err := m.DeleteSystemInfo("1234567890")
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestSysInfo_RetrieveSystemInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockSysInfoer(ctrl)

	m.EXPECT().RetrieveSystemInfo(gomock.Any()).Return(nil, nil)
	convey.Convey("TestSysInfo_RetrieveSystemInfo", t, func() {
		val, err := m.RetrieveSystemInfo(bson.M{})
		convey.So(val, convey.ShouldEqual, nil)
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestSysInfo_UpdateSystemInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockSysInfoer(ctrl)

	m.EXPECT().UpdateSystemInfo(gomock.Any()).Return(nil)
	convey.Convey("TestSysInfo_DeleteSystemInfo", t, func() {
		err := m.UpdateSystemInfo(&SysInfo{
			SystemInfo: SystemInfo{},
			CreateTime: primitive.Timestamp{},
		})
		convey.So(err, convey.ShouldEqual, nil)
	})
}
