package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	. "monitor_server/utils"
)

type AppCode struct {
	AppID   string `json:"app_id"`
	AppCode string `json:"app_code"`
}

const (
	KMongoAppCodeCollection = "app_code"
)

func RetrieveAppCode(filter bson.M) (*AppCode, error) {
	var appCode *AppCode
	res := GlobalDatabase.Collection(KMongoAppCodeCollection).FindOne(context.TODO(), filter)
	err := res.Decode(&appCode)
	if err != nil {
		if err == mongo.ErrNilDocument || err == mongo.ErrNoDocuments {
			return nil, nil
		}
		SugarLogger.Error("MONGODB ERROR@RetrieveAppCode, Error Info:", err)
		return nil, err
	}
	return appCode, nil
}
