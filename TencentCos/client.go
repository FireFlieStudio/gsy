package TencentCos

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

func NewClient(BucketUrl string) *cos.Client {
	u, _ := url.Parse(BucketUrl)
	return cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretId,
			SecretKey: secretKey,
		},
	})
}

func DefaultClient() *cos.Client {
	return NewClient(defaultBucketUrl)
}
