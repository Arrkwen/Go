/*
 * @Author: your name
 * @Date: 2022-04-16 20:23:42
 * @LastEditTime: 2022-04-17 15:14:47
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/cmd/userInfoManageServer/main.go
 */

package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Arrkwen/Go/userInfo/api"
	"github.com/Arrkwen/Go/userInfo/service"
	"github.com/Arrkwen/Go/userInfo/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	configPath = flag.String("config", "config/uims_config.json", "user information managerment server config")
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
 * @description: 初始化服务对象
 * @param {*utils.ServerConfig} cfg：服务配置选项
 * @return {*}：返回服务对象指针
 */
func initUIMS(cfg *utils.ServerConfig) *service.UserInfoManagerServer {
	server := service.NewUserInfoManagerServer(cfg)
	return server
}

func main() {
	// 加载配置，并初始化服务对象和日志
	flag.Parse()
	cfg := utils.ServerConfig{}
	utils.LoadConfig(*configPath, &cfg)
	initLogger(cfg.LogLevel)
	uims := initUIMS(&cfg)

	// 注册 gprc服务
	listener, err := net.Listen("tcp", cfg.GPRCEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcOption := utils.ServerOptionGRPC(&cfg)
	grpcServer := grpc.NewServer(grpcOption...)
	api.RegisterUserServiceServer(grpcServer, uims)

	// 启动 grpc 服务
	go func() {

		if err := grpcServer.Serve(listener); err != nil {
			log.WithError(err).Fatal("failed to serve grpc service")
		}
	}()
	log.Infof("Server start.....")
	log.Infof("Listen %v", cfg.GPRCEndpoint)

	// 停止 grpc 服务
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	log.Infof("user info management server existing...")
	grpcServer.GracefulStop()
	log.Infof("user info management server has exited!")
}
