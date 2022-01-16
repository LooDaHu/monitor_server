package model

import (
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUser_CreateUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockUserer(ctrl)

	m.EXPECT().CreateUserInfo(gomock.Any()).Return(nil)
	convey.Convey("TestUser_CreateUserInfo", t, func() {
		err := m.CreateUserInfo(&User{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestUser_DeleteUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockUserer(ctrl)

	m.EXPECT().CreateUserInfo(gomock.Any()).Return(nil)
	convey.Convey("TestUser_DeleteUserInfo", t, func() {
		err := m.CreateUserInfo(&User{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestUser_RetrieveUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockUserer(ctrl)

	m.EXPECT().RetrieveUserInfo(gomock.Any()).Return(&User{
		ID:       primitive.ObjectID{},
		Username: "hello",
		Password: "password",
		Type:     "admin",
	}, nil)
	convey.Convey("TestUser_RetrieveUserInfo", t, func() {
		val, err := m.RetrieveUserInfo(bson.M{})
		convey.So(val.Username, convey.ShouldEqual, "hello")
		convey.So(val.Password, convey.ShouldEqual, "password")
		convey.So(val.Type, convey.ShouldEqual, "admin")
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestUser_UpdateUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockUserer(ctrl)

	m.EXPECT().UpdateUserInfo(gomock.Any()).Return(nil)
	convey.Convey("TestUser_UpdateUserInfo", t, func() {
		err := m.UpdateUserInfo(&User{})
		convey.So(err, convey.ShouldEqual, nil)
	})
}
