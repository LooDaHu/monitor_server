package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"

	. "monitor_server/utils"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`           // 订单id
	Username string             `bson:"username" json:"username"` //用户名称
	Password string             `bson:"password" json:"password"` //用户密码
	Type     string             `bson:"type" json:"type"`         // 用户类型
}

type Session struct {
	User string `json:"user"`
}

const (
	KMongoUserCollection = "User"
	KUserTypeAdmin       = "admin"
	KUserTypeGuest       = "guest"
)

// CheckUserInfo 检查用户信息
func CheckUserInfo(username string, password string) (bool, string, string, error) {
	var user User
	filter := bson.M{
		"username": username,
		"password": password,
	}
	res := GlobalDatabase.Collection(KMongoUserCollection).FindOne(context.TODO(), filter)
	err := res.Decode(&user)
	if err != nil {
		if err == mongo.ErrNilDocument || err == mongo.ErrNoDocuments {
			return false, "", "", nil
		}
		SugarLogger.Error("MONGODB ERROR@CheckUserInfo, Error Info:", err)
		return false, "", "", err
	}
	return true, user.ID.Hex(), user.Type, nil
}
