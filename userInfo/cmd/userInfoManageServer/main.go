/*
 * @Author: your name
 * @Date: 2022-04-16 20:23:42
 * @LastEditTime: 2022-04-17 12:52:59
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/cmd/userInfoManageServer/main.go
 */

package main

import (
	"flag"
	"net"

	"github.com/Arrkwen/Go/userInfo/api"
	"github.com/Arrkwen/Go/userInfo/service"
	"github.com/Arrkwen/Go/userInfo/utils"
	"github.com/prometheus/common/log"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	configPath = flag.String("config", "config/uims_config.json", "user information managerment server config")
)

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

// initServer
func initUIMS(cfg *utils.ServerConfig) *service.UserInfoManagerServer {
	server := service.NewUserInfoManagerServer(cfg)
	return server
}

func main() {
	flag.Parse()
	cfg := utils.ServerConfig{}
	utils.LoadConfig(*configPath, &cfg)
	initLogger(cfg.LogLevel)
	uims := initUIMS(cfg)

	// 注册gprc服务
	listener, err := net.Listen("tcp", cfg.GPRCEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcOption := utils.ServerOptionGRPC(cfg)
	grpcServer := grpc.NewServer(grpcOption...)
	api.RegisterUserServiceServer(grpcServer, uims)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.WithError(err).Fatal("failed to serve grpc service")
		}
		log.Infof("Server start.....")
	}()

	// 注册http 服务

}
