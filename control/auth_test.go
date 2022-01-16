package control

import (
	"github.com/gin-gonic/gin"
	"monitor_server/message"
	"monitor_server/model"
	"testing"
)

func TestCheckUserInfo(t *testing.T) {
	type args struct {
		req  *message.LoginReq
		user *model.User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUserInfo(tt.args.req, tt.args.user); got != tt.want {
				t.Errorf("CheckUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestControllerAuth_AppCodeCheckRequired(t *testing.T) {
	type fields struct {
		user    model.User
		appCode model.AppCode
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := ControllerAuth{
				user:    tt.fields.user,
				appCode: tt.fields.appCode,
			}
			ca.AppCodeCheckRequired(tt.args.c)
		})
	}
}

func TestControllerAuth_Login(t *testing.T) {
	type fields struct {
		user    model.User
		appCode model.AppCode
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := ControllerAuth{
				user:    tt.fields.user,
				appCode: tt.fields.appCode,
			}
			ca.Login(tt.args.c)
		})
	}
}

func TestControllerAuth_LoginRequired(t *testing.T) {
	type fields struct {
		user    model.User
		appCode model.AppCode
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := ControllerAuth{
				user:    tt.fields.user,
				appCode: tt.fields.appCode,
			}
			ca.LoginRequired(tt.args.c)
		})
	}
}

func TestControllerAuth_Logout(t *testing.T) {
	type fields struct {
		user    model.User
		appCode model.AppCode
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ca := ControllerAuth{
				user:    tt.fields.user,
				appCode: tt.fields.appCode,
			}
			ca.Logout(tt.args.c)
		})
	}
}
