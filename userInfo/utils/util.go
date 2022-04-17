/*
 * @Author: your name
 * @Date: 2022-04-16 20:30:34
 * @LastEditTime: 2022-04-17 14:17:27
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/utils/util.go
 */

package utils

import (
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/keepalive"
)

/**
 * @description: 初始化gPRC服务端配置选项
 * @param {*ServerConfig} cfg：配置结构体输入
 * @return {*}：服务端配置对象
 */
func ServerOptionGRPC(cfg *ServerConfig) []grpc.ServerOption {
	var opts []grpc.ServerOption
	if cfg.TLS {
		if cfg.CertFilePath == "" {
			cfg.CertFilePath = data.Path("x509/server_cert.pem")
		}
		if cfg.KeyFilePath == "" {
			cfg.KeyFilePath = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(cfg.CertFilePath, cfg.KeyFilePath)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    time.Minute,
		Timeout: time.Second * 30,
	}))

	opts = append(opts, grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             1 * time.Minute,
		PermitWithoutStream: true,
	}))

	const minMsgSize = 4
	const maxMsgSize = 64
	if cfg.MsgSize > maxMsgSize {
		cfg.MsgSize = maxMsgSize
		log.Warnf("gRPC max message size should not be larger than 64MB, actual size: %d, change to 64MB", maxMsgSize)
	}

	if cfg.MsgSize < minMsgSize {
		cfg.MsgSize = minMsgSize
		log.Warnf("gRPC max message size should be larger than 4MB, actual size: %d, change to 4MB", minMsgSize)
	}
	opts = append(opts, grpc.MaxSendMsgSize(cfg.MsgSize*1024*1024))
	opts = append(opts, grpc.MaxRecvMsgSize(cfg.MsgSize*1024*1024))

	return opts
}

func ClientOptionsGRPC(cfg *ClientConfig) []grpc.DialOption {
	var opts []grpc.DialOption
	if cfg.TLS {
		if cfg.CaFilePath == "" {
			cfg.CaFilePath = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(cfg.CaFilePath, cfg.ServerHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	return opts
}
