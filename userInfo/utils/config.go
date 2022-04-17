/*
 * @Author: your name
 * @Date: 2022-04-16 20:31:49
 * @LastEditTime: 2022-04-17 12:37:47
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /testCode/userInfo/utils/config.go
 */

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
)

type ServerConfig struct {
	ServerName   string
	GPRCEndpoint string
	HTTPEndpoint string
	TLS          bool
	CertFilePath string
	KeyFilePath  string
	MsgSize      int32
	LogLevel     string
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

func LoadConfig(configPath string, ptr interface{}) {
	if err := loadConfig(configPath, ptr); err != nil {
		log.Fatal(err)
	}
}
