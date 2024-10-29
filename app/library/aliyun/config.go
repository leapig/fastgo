package aliyun

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
)

type AliYun struct {
	AccessKeyId     string
	AccessKeySecret string
	OssRegion       string
	OssBucketName   string
	OssClient       *oss.Client
}
