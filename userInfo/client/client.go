/*
 * @Author: your name
 * @Date: 2022-04-16 20:21:53
 * @LastEditTime: 2022-04-17 15:47:30
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/client/client.go
 */

package main

import (
	"context"
	"flag"
	"fmt"

	"time"

	"github.com/Arrkwen/Go/userInfo/api"
	"github.com/Arrkwen/Go/userInfo/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	configPath = flag.String("config", "client_config.json", "client configuration")
)

/**
 * @description: 初始化日志
 * @param {string} level：trace,debug,info,warning,error,fatal,panic
 * @return {*}
 */
func initLogger(level string) {
	var logLevel log.Level
	var err error
	if level != "" {
		logLevel, err = log.ParseLevel(level)
		if err != nil {
			logLevel = log.InfoLevel
		}
	} else {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
}

/**
 * @description: 客户端接口实现
 * @param {api.UserServiceClient} client：grpc stub client
 * @param {*api.User} userInfo: 需要保存的用户信息
 * @return {*}
 */
func SaveUserInfo(client api.UserServiceClient, userInfo *api.User) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rsp, err := client.SaveUserInfo(ctx, &api.SaveUserInfoRequest{User: userInfo})
	if err != nil {
		return false
	}
	return rsp.GetIsSuccess()
}

func main() {
	// 加载配置,并初始化日志
	flag.Parse()
	cfg := utils.ClientConfig{}
	err := utils.LoadConfig(*configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	initLogger(cfg.LogLevel)
	log.Infof("connecting : %v", cfg.GPRCEndpoint)

	// 连接服务端
	clientOption := utils.ClientOptionsGRPC(&cfg)
	conn, err := grpc.Dial(cfg.GPRCEndpoint, clientOption...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := api.NewUserServiceClient(conn)
	log.Infof("connect success: %v", client)

	// 构造请求数据
	localTime := time.Now()
	reqData := api.User{
		UserId:           "0",
		UserImage:        "dijia",
		UserPhone:        "155xxxx8388",
		UserPassword:     "*******",
		UserType:         1,
		UserRegisterTime: timestamppb.New(localTime),
		UserSchool:       "USTC",
		UserResearch:     "CV",
		UserGithub:       "github.com/xxx",
		UserGoogle:       "https://scholar.google.com/xxx",
	}
	// 服务调用
	isSuccess := SaveUserInfo(client, &reqData)
	if isSuccess {
		fmt.Println("Save user info success")
	} else {
		fmt.Println("Save user info failed")
	}
}
