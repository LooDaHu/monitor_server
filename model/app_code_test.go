package model

import (
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestRetrieveAppCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockCoder(ctrl)

	m.EXPECT().RetrieveAppCode(gomock.Not(bson.M{
		"app_id": KTestAppId,
	})).Return(nil, mongo.ErrNoDocuments)
	m.EXPECT().RetrieveAppCode(gomock.Eq(bson.M{
		"app_id": KTestAppId,
	})).Return(&AppCode{
		AppID:   KTestAppId,
		AppCode: KTestAppCode,
	}, nil)
	convey.Convey("Test Retrieve Right App Code", t, func() {
		appCode, err := m.RetrieveAppCode(bson.M{
			"app_id": KTestAppId,
		})
		convey.So(appCode.AppID, convey.ShouldEqual, KTestAppId)
		convey.So(appCode.AppCode, convey.ShouldEqual, KTestAppCode)
		convey.So(err, convey.ShouldEqual, nil)
	})

	convey.Convey("Test Retrieve Wrong App Code", t, func() {
		appCode, err := m.RetrieveAppCode(bson.M{
			"app_id": "random_wrong_code",
		})
		convey.So(appCode, convey.ShouldEqual, nil)
		convey.So(err, convey.ShouldEqual, mongo.ErrNoDocuments)
	})
}
