/*
 * @Author: your name
 * @Date: 2022-04-16 20:01:39
 * @LastEditTime: 2022-04-17 12:52:04
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/server/service.go
 */
package service

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Arrkwen/Go/userInfo/api"
	"github.com/Arrkwen/Go/userInfo/utils"
)

/**
 * @description: UserManagerrServer is the API for  user information management service
 * @attribute {string} name : server name
 * @attribute {UnimplementedRouteGuideServer}: it must add into the server struct
 */
type UserInfoManagerServer struct {
	name string
	api.UnimplementedRouteGuideServer
}

/**
 * @description: New a server
 * @param {utils.ServerConfig} cfg: server configuration
 * @return {*UserInfoManagerServer}: server object pointer
 */
func NewUserInfoManagerServer(cfg utils.ServerConfig) *UserInfoManagerServer {
	server := &UserInfoManagerServer{name: cfg.ServerName}
	return server
}

/**
 * @description: rpc api:SaveUserInfo
 * @param {context.Context} ctx
 * @param {*api.SaveUserInfoRequest} req: rpc request
 * @return {api.SaveUserInfoResponse}rsp: rpc response
 */
func (u *UserInfoManagerServer) SaveUserInfo(ctx context.Context, req *api.SaveUserInfoRequest) (*api.SaveUserInfoResponse, error) {
	v := reflect.ValueOf(req.User)
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
		case reflect.Int32:
			fmt.Println(f.Int())
		}
	}
	return api.SaveUserInfoResponse{true}, nil
}
