/*
 * @Author: your name
 * @Date: 2022-04-16 20:01:39
 * @LastEditTime: 2022-04-17 15:41:58
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/server/service.go
 */
package service

import (
	"context"
	"fmt"

	"github.com/Arrkwen/Go/userInfo/api"
	"github.com/Arrkwen/Go/userInfo/utils"
	log "github.com/sirupsen/logrus"
)

/**
 * @description: UserManagerrServer is the API for  user information management service
 * @attribute {string} name : server name
 * @attribute {UnimplementedRouteGuideServer}: it must add into the server struct
 */
type UserInfoManagerServer struct {
	name string
	api.UnimplementedUserServiceServer
}

/**
 * @description: New a server
 * @param {utils.ServerConfig} cfg: server configuration
 * @return {*UserInfoManagerServer}: server object pointer
 */
func NewUserInfoManagerServer(cfg *utils.ServerConfig) *UserInfoManagerServer {
	server := &UserInfoManagerServer{name: cfg.ServerName}
	return server
}

/**
 * @description: rpc api:SaveUserInfo:暂时实现是打印请求数据
 * @param {context.Context} ctx：上下文，暂时未使用
 * @param {*api.SaveUserInfoRequest} req: rpc request：请求数据
 * @return {api.SaveUserInfoResponse}rsp: rpc response：响应数据
 */
func (u *UserInfoManagerServer) SaveUserInfo(ctx context.Context, req *api.SaveUserInfoRequest) (*api.SaveUserInfoResponse, error) {
	log.Infof("Saving user info...")
	fmt.Println("%v", req.User)
	return &api.SaveUserInfoResponse{IsSuccess: true}, nil
}
