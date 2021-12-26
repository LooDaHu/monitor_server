package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "monitor_server/utils"
)

const (
	KMongoSysInfoCollection = "sys_info"
)

type SysInfo struct {
	SystemInfo SystemInfo          `json:"system_info"`
	CreateTime primitive.Timestamp `json:"create_time"`
}

type SystemInfo struct {
	Host    HostInfo    `json:"host"`
	Network NetworkInfo `json:"network"`
	CPU     CPUInfo     `json:"cpu"`
	Memory  MemoryInfo  `json:"memory"`
	Disk    DiskInfo    `json:"disk"`
}

type HostInfo struct {
	HostName string `json:"host_name"`
	OS       string `json:"os"`
}

type NetworkInfo struct {
	IPAddress   string `json:"ip_address"`
	BytesSend   uint64 `json:"bytes_send"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
}

type CPUInfo struct {
	ModelName string  `json:"model_name"`
	Logical   int     `json:"logical"`
	Physical  int     `json:"physical"`
	Percent   float64 `json:"percent"`
}

type MemoryInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Swap  uint64 `json:"swap"`
}

type DiskInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

type SysInfoer interface {
	CreateSystemInfo(sysInfo *SysInfo) error
	RetrieveSystemInfo(filter bson.M) (*SysInfo, error)
	UpdateSystemInfo(sysInfo *SysInfo) error
	DeleteSystemInfo(id string) error
}

func (s SysInfo) UpdateSystemInfo(sysInfo *SysInfo) error {
	return nil
}

func (s SysInfo) DeleteSystemInfo(id string) error {
	return nil
}

func (s SysInfo) CreateSystemInfo(sysInfo *SysInfo) error {
	_, err := GlobalDatabase.Collection(
		KMongoSysInfoCollection).InsertOne(context.TODO(), sysInfo)
	if err != nil {
		return err
	}
	return nil
}

func (s SysInfo) RetrieveSystemInfo(filter bson.M) (*SysInfo, error) {
	var info *SysInfo
	res := GlobalDatabase.Collection(KMongoSysInfoCollection).FindOne(context.TODO(), filter)
	err := res.Decode(&info)
	if err != nil {
		if err == mongo.ErrNilDocument || err == mongo.ErrNoDocuments {
			return nil, nil
		}
		SugarLogger.Error("MONGODB ERROR@RetrieveSystemInfo, Error Info:", err)
		return nil, err
	}
	return info, nil
}
