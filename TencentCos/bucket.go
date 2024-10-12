package TencentCos

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gsync/logger"
	"strings"
)

func CreateBucket(bucketName, regionKey string) {
	bucketUrl := GenerateBucketUrl(bucketName, regionKey)
	client := NewClient(bucketUrl)
	opt := &cos.BucketPutOptions{
		XCosACL: "private",
	}

	_, err := client.Bucket.Put(context.Background(), opt)
	if err != nil {
		logger.Error("创建bucket失败[ %s ] %s", bucketUrl, err.Error())
	} else {
		logger.Info("创建bucket成功[ %s ] ", bucketUrl)
	}

}

func DeleteBucket(bucketName string) {
	bucketUrl, ok := FindBucketUrlByName(bucketName)
	if !ok {
		logger.Error("bucket不存在[ %s ],无法删除", bucketName)
	}
	client := NewClient(bucketUrl)
	_, err := client.Bucket.Delete(context.Background())
	if err != nil {
		logger.Error("删除bucket失败[ %s ] %s", bucketUrl, err.Error())
	} else {
		logger.Info("删除bucket成功[ %s ] ", bucketUrl)
	}
}

func ListBuckets() []string {
	client := DefaultClient()
	s, _, err := client.Service.Get(context.Background())
	if err != nil {
		logger.Error("获取Bucket错误 %s", err.Error())
		return nil
	}

	var buckets []string
	for _, b := range s.Buckets {
		bUrl := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", b.Name, b.Region)
		buckets = append(buckets, bUrl)
	}
	return buckets
}

func FindBucketUrlByName(bucketName string) (string, bool) {
	buckets := ListBuckets()
	for _, bucket := range buckets {
		bName, _ := GetBucketName(bucket)
		if bucketName == bName {
			return bucket, true
		}
	}
	return "", false
}

func IsBucketExist(bucketUrl string) bool {
	client := NewClient(bucketUrl)
	ok, err := client.Bucket.IsExist(context.Background())
	if err == nil && ok {
		return true
	}
	return false
}

func GenerateBucketUrl(bucketName, regionKey string) string {
	regions := map[string]string{"北京": "beijing", "南京": "nanjing", "上海": "shanghai", "广州": "guangzhou", "成都": "chengdu", "重庆": "chongqing"}

	bUrl, err := cos.NewBucketURL(fmt.Sprintf("%s-%s", bucketName, appId), fmt.Sprintf("ap-%s", regions[regionKey]), true)
	if err != nil {
		logger.Error("生成BucketUrl错误: %s", err.Error())
		return ""
	}
	return bUrl.String()
}

func GetBucketName(bucketUrl string) (string, string) {
	bucketId := strings.Split(bucketUrl, ".")[0]
	bucketId = strings.Replace(bucketId, "https://", "", -1)
	bucketId = strings.Replace(bucketId, "http://", "", -1)

	bucketName := strings.Split(bucketId, "-")[0]
	appId := strings.Split(bucketId, "-")[1]
	return bucketName, appId
}
