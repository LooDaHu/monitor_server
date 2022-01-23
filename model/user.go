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

type Userer interface {
	CreateUserInfo(user *User) error
	RetrieveUserInfo(filter bson.M) (*User, error)
	UpdateUserInfo(user *User) error
	DeleteUserInfo(id string) error
}

func (u User) CreateUserInfo(user *User) error {
	return nil
}

func (u User) UpdateUserInfo(user *User) error {
	return nil
}

func (u User) DeleteUserInfo(id string) error {
	return nil
}

// RetrieveUserInfo 查询用户信息
func (u User) RetrieveUserInfo(filter bson.M) (*User, error) {
	var user *User
	res := GlobalDatabase.Collection(KMongoUserCollection).FindOne(context.TODO(), filter)
	err := res.Decode(&user)
	if err != nil {
		if err == mongo.ErrNilDocument || err == mongo.ErrNoDocuments {
			return &User{}, nil
		}
		SugarLogger.Error("MONGODB ERROR@RetrieveUserInfo, Error Info:", err)
		return nil, err
	}
	return user, nil
}
