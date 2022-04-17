/*
 * @Author: your name
 * @Date: 2022-04-16 20:31:49
 * @LastEditTime: 2022-04-17 14:36:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/utils/config.go
 */

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
)

// 服务端配置
type ServerConfig struct {
	ServerName   string // 服务名
	GPRCEndpoint string // grpc服务端的host:port
	HTTPEndpoint string // HTTP服务的host:port
	TLS          bool   // 是否TLS加密
	CertFilePath string // 加密证书路径
	KeyFilePath  string // 秘钥路径
	MsgSize      int    // gRPC收发包大小限制[4-64]MB
	LogLevel     string // 日志等级：trace,debug,info,warning,error,fatal,panic
}

// 客户端配置
type ClientConfig struct {
	GPRCEndpoint       string // grpc服务端的host:port
	HTTPEndpoint       string // HTTP服务的host:port
	ServerHostOverride string // gRPC服务端域名："xx.xxx.com"
	TLS                bool   // 是否TLS加密
	CaFilePath         string // 证书路径
	LogLevel           string // 日志等级：trace,debug,info,warning,error,fatal,panic
}

func loadConfig(configPath string, ptr interface{}) error {
	if ptr == nil {
		return fmt.Errorf("ptr of type(%T) is nil", ptr)
	}
	grpc.WithBlock()

	data, err := ioutil.ReadFile(configPath) // #nosec
	if err != nil {
		return fmt.Errorf("open file(%v) with err %v", configPath, err)
	}

	if err := json.Unmarshal(data, ptr); err != nil {
		return err
	}

	return nil
}

/**
 * @description: 根据配置文件路径解码到对应的结构体
 * @param {string} configPath：配置路径
 * @param {interface{}} ptr：解码对象
 * @return {*}：是否解码成功
 */
func LoadConfig(configPath string, ptr interface{}) error {
	if err := loadConfig(configPath, ptr); err != nil {
		return err
	}
	return nil
}
