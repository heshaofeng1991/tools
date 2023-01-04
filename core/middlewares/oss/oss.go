package oss

import (
	"core/config"
	"io"
	"net/url"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Bucket struct {
	*oss.Bucket
	err error
}

func New(conf config.Oss) (bucket *Bucket) {
	bucket = &Bucket{}
	client, err := oss.New(conf.Endpoint, conf.Id, conf.Secret)
	if err != nil {
		bucket.err = err
		return
	}
	bucket.Bucket, bucket.err = client.Bucket(conf.Bucket)
	return
}

func (bucket *Bucket) bucketName(bucketName string) *Bucket {
	if bucket.err != nil {
		return bucket
	}
	bucket.Bucket, bucket.err = bucket.Client.Bucket(bucketName)
	return bucket
}

func (bucket *Bucket) Upload(filePath string, reader io.Reader) *Bucket {
	if bucket.err != nil {
		return bucket
	}
	bucket.err = bucket.PutObject(filePath, reader)
	return bucket
}

func (bucket *Bucket) Del(filePath string) *Bucket {
	if bucket.err != nil {
		return bucket
	}
	bucket.err = bucket.DeleteObject(filePath)
	return bucket
}

func (bucket *Bucket) Err() error {
	return bucket.err
}

func (bucket *Bucket) GetOssUrl(filePath string) (string, error) {
	if bucket.err != nil {
		return "", nil
	}
	parse, _ := url.Parse(bucket.GetConfig().Endpoint)
	return bucket.BucketName + "." + parse.Host + "/" + filePath, nil
}

func (bucket *Bucket) Download(filePath string) (string, error) {
	return bucket.GetOssUrl(filePath)
}
