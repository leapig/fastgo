package aliyun

import (
	"bytes"
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/leapig/fastgo/app/dal/model"
	"os"
)

// Oss 实例对象
func Oss() *AliYun {
	aliYunOss := &AliYun{
		AccessKeyId:     os.Getenv("ALIYUN_ACCESSKEYID"),
		AccessKeySecret: os.Getenv("ALIYUN_ACCESSKEYSECRET"),
		OssRegion:       os.Getenv("ALIYUN_OSS_REGION"),
		OssBucketName:   os.Getenv("ALIYUN_OSS_BUCKETNAME"),
	}
	aliYunOss.ossConnect()
	return aliYunOss
}

// ossConnect 连接阿里云oss
func (a *AliYun) ossConnect() {
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(a.AccessKeyId, a.AccessKeySecret)).WithRegion(a.OssRegion)
	a.OssClient = oss.NewClient(cfg)
	_, err := a.OssClient.GetBucketInfo(context.TODO(), &oss.GetBucketInfoRequest{
		Bucket: &a.OssBucketName,
	})
	if err != nil {
		panic(err)
	}
}

// PutObject 上传对象文件
func (a *AliYun) PutObject(object model.Oss) (model.Oss, error) {
	u := a.OssClient.NewUploader()
	_, err := u.UploadFrom(context.TODO(), &oss.PutObjectRequest{
		Bucket: &a.OssBucketName,
		Key:    &object.Name,
	}, bytes.NewBuffer(object.Data))
	if err == nil {
		return a.PreSignedUrlForGet(object)
	}
	return object, err
}

// PreSignedUrlForGet 获取只读授权链接
func (a *AliYun) PreSignedUrlForGet(object model.Oss) (model.Oss, error) {
	result, err := a.OssClient.Presign(context.TODO(), &oss.GetObjectRequest{
		Bucket: &a.OssBucketName,
		Key:    &object.Name,
	})
	if err == nil {
		object.Url = result.URL
	}
	return object, err
}

// PathToUrl 名称换链接
func (a *AliYun) PathToUrl(path string) string {
	if picUrl, err := a.PreSignedUrlForGet(model.Oss{
		Name: path,
	}); err == nil {
		return picUrl.Url
	} else {
		return path
	}
}

// RemoveObject 移除对象文件
func (a *AliYun) RemoveObject(object model.Oss) error {
	_, err := a.OssClient.DeleteObject(context.TODO(), &oss.DeleteObjectRequest{
		Bucket: &a.OssBucketName,
		Key:    &object.Name,
	})
	return err
}
