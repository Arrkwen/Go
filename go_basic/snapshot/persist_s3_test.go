/*
 * @Author: xiaokun xiaokun@sensetime.com
 * @Date: 2022-05-16 13:45:46
 * @LastEditors: xiaokun xiaokun@sensetime.com
 * @LastEditTime: 2022-05-16 16:42:52
 * @FilePath: /Go/snapshot/persist_s3_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package snapshot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS3Persist(t *testing.T) {
	s3config := &S3Config{
		Endpoint:          "10.151.5.121:30021",
		AccessKeyID:       "minio",
		SecretAccessKeyID: "qqnhEylFQZbvmlV",
		UseSSL:            false,
		BucketName:        "rtc",
		Location:          "cn-north-1",
	}
	// V1 老版本代码测试
	s3V1persist, err := NewS3PersistV1(s3config)
	s3V1persist.Init()
	assert.NoError(t, err)
	testBytes := []byte("tets")
	err = s3V1persist.WriteMeta(1, testBytes)
	assert.NoError(t, err)
	res, err := s3V1persist.ReadMeta(1)
	assert.NoError(t, err)
	fmt.Printf("res = %v\n", res)
	assert.NoError(t, err)
	err = s3V1persist.DeleteMeta(1)
	assert.NoError(t, err)
	// V7 新版本代码测试
	s3V7persist, err := NewS3PersistV7(s3config)
	s3V7persist.Init()
	assert.NoError(t, err)
	err = s3V7persist.WriteMeta()
	assert.NoError(t, err)
	meta := &Meta{}
	s3V7persist.ReadMeta(meta)
	fmt.Printf("meta = %v\n", meta)
	s3V7persist.DeleteMeta(meta)

}
