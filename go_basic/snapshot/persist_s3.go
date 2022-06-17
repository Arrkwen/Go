/*
 * @Author: your name
 * @Date: 2022-04-19 14:06:41
 * @LastEditTime: 2022-05-16 16:36:25
 * @LastEditors: xiaokun xiaokun@sensetime.com
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /realtime-clustering/snapshot/persist_s3.go
 */
package snapshot

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go"
	minioV1 "github.com/minio/minio-go"
	minioV7 "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

type S3Config struct {
	Endpoint          string
	AccessKeyID       string
	SecretAccessKeyID string
	UseSSL            bool
	BucketName        string
	Location          string // timeout(ms) with connection to S3 persistent
}

type Meta struct {
	tmpPath  string // 临时路径
	savePath string // 保存路径
	data     []byte // 序列化数据
}

type S3PersistV1 struct {
	prefix     string
	bucketName string
	location   string
	timeout    int
	client     *minioV1.Client
}

type S3PersistV7 struct {
	prefix     string
	bucketName string
	location   string
	timeout    int
	client     *minioV7.Client
}

const (
	TIME_OUT     = 6
	RETRY_COUNTS = 3
)

func NewS3PersistV1(config *S3Config) (*S3PersistV1, error) {
	client, err := minioV1.New(config.Endpoint, config.AccessKeyID, config.SecretAccessKeyID, config.UseSSL)
	if err != nil {
		return nil, err
	}

	dir := "tes-snapshots"
	if err != nil {
		return nil, err
	}
	log.Infof("connect s3 sucessfully")
	return &S3PersistV1{
		prefix:     dir,
		bucketName: config.BucketName,
		location:   config.Location,
		client:     client,
		timeout:    TIME_OUT,
	}, nil
}

func NewS3PersistV7(config *S3Config) (*S3PersistV7, error) {
	rootPath := "tes-snapshots"
	opts := &minioV7.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKeyID, ""),
		Secure: config.UseSSL,
	}
	client, err := minioV7.New(config.Endpoint, opts)
	if err != nil {
		return nil, err
	}
	log.Debug("S3Persist client created successfully")
	if strings.Contains(rootPath, ".") || strings.Contains(rootPath, "..") {
		rootPath = strings.TrimPrefix(rootPath, "./")
		rootPath = strings.TrimPrefix(rootPath, "../")
	}
	return &S3PersistV7{
		prefix:     rootPath,
		bucketName: config.BucketName,
		location:   config.Location,
		client:     client,
		timeout:    TIME_OUT,
	}, nil
}

func (s *S3PersistV1) metaFullPath(seq int64) string {
	return filepath.Join(s.prefix, fmt.Sprintf("META-%016d.json", seq))
}
func (s *S3PersistV1) Init() error {
	ok, err := s.client.BucketExists(s.bucketName)
	if err != nil {
		return err
	}
	if !ok {
		if err := s.client.MakeBucket(s.bucketName, s.location); err != nil {
			return err
		}
	}
	return nil
}
func (s *S3PersistV1) ReadMeta(seq int64) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.timeout)*time.Second)
	defer cancel()
	obj, err := s.client.GetObjectWithContext(ctx, s.bucketName, s.metaFullPath(seq), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := obj.Close(); err != nil {
			log.Errorf("failed to close Object: %s", err)
		}
	}()
	return ioutil.ReadAll(obj)
}

func (s *S3PersistV1) WriteMeta(seq int64, b []byte) error {
	buf := bytes.NewBuffer(b)
	flen := int64(len(b))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.timeout)*time.Second)
	defer cancel()
	savePath := s.metaFullPath(seq)
	log.Infof("writing meta %v", savePath)
	n, err := s.client.PutObjectWithContext(ctx, s.bucketName, s.metaFullPath(seq), buf, flen, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	if n != flen {
		return errors.New("partial write")
	}
	log.Infof("write Successfully")
	return nil
}

func (s *S3PersistV1) DeleteMeta(seq int64) error {
	p := s.metaFullPath(seq)
	return s.client.RemoveObject(s.bucketName, p)
}

func (s *S3PersistV7) Init() error {
	ctx := context.Background()
	err := s.client.MakeBucket(ctx, s.bucketName, minioV7.MakeBucketOptions{Region: s.location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := s.client.BucketExists(ctx, s.bucketName)
		if errBucketExists == nil && exists {
			log.Debugf("the bucket %s already exists", s.bucketName)
		} else {
			log.WithError(err).Error("Faild create bucket!")
			return err
		}
	} else {
		log.Debugf("Successfully created %s\n", s.bucketName)
	}
	if err := os.MkdirAll(s.prefix, 0750); err != nil {
		return err
	}
	return nil
}

func (s *S3PersistV7) WriteMeta() error {
	file, err := os.Open("file")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return err
	}

	uploadInfo, err := s.client.PutObject(context.Background(), s.bucketName, "myobject/file", file, fileStat.Size(), minioV7.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
	return nil
}

func (s *S3PersistV7) ReadMeta(meta *Meta) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.timeout)*time.Second)
	defer cancel()
	var err error
	object, err := s.client.GetObject(ctx, s.bucketName, "myobject/file", minioV7.GetObjectOptions{})
	if err != nil {
		log.Errorf("ReadSnapshot: Failed to get object err: %v", err)
		return err
	}
	defer func() {
		if err := object.Close(); err != nil {
			log.Errorf("ReadSnapshot failed to close Object: %s", err)
		}
	}()
	meta.data, err = ioutil.ReadAll(object)
	if err != nil {
		log.Errorf("ReadSnapshot read snapshot from Object error: %v", err)
		return err
	}

	log.Infof("ReadSnapshot read snapshot from s3 successfully<%s>", meta.savePath)
	return nil
}

func (s *S3PersistV7) DeleteMeta(meta *Meta) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.timeout)*time.Second)
	defer cancel()

	opts := minioV7.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	for i := 0; i < RETRY_COUNTS; i++ {
		err = s.client.RemoveObject(ctx, s.bucketName, "myobject/file", opts)
		if err == nil {
			break
		}
		log.Warnf("failed to delete s3 snapshot<%s> , retry=%d", "myobject/file", i)
		time.Sleep(3 * time.Second)
	}
	log.Debugf("delete the snapshot <%s> at s3", meta.savePath)
	return err
}
