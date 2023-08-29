package oss

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"tiktok_demo/pkg/constants"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
    client *oss.Client
    err error
)

// CreateBucket create a bucket with a specified name
func CreateBucket(bucketName string) {
    exists, err := client.IsBucketExist(bucketName)
    if err != nil {
        fmt.Println(err)
        return
    }
    if !exists {
        err = client.CreateBucket(bucketName)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Successfully created mybucket %v\n", bucketName)
    }
}

func PutToBucket(bucketName string, file *multipart.FileHeader) (err error) {
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        return err
    }
    fileObj, _ := file.Open()
    defer fileObj.Close()
    err = bucket.PutObject(bucketName + file.Filename, fileObj)
    return err
}

// GetObjURL get the original link of the file in oss
func GetObjURL(bucketName, objectName string, expTime int64) (u string, err error) {
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        return "", err
    }
    u, err = bucket.SignURL(bucketName + objectName, oss.HTTPGet, expTime)
    return u, err
}

func PutToBucketByBuf(bucketName, objectName string, buf *bytes.Buffer) (err error) {
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        return err
    }
    err = bucket.PutObject(bucketName + objectName, buf)
    return
}

func PutToBucketByFilePath(bucketName, filename, filepath string) (err error) {
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        return err
    }
    err = bucket.PutObjectFromFile(bucketName + filename, filepath)
    return err
}

func Init() {
    client, err = oss.New(constants.OSSEndPoint, constants.OSSAccessKeyID, constants.OSSAccessKeySecret)
    if err != nil {
        log.Fatalln("OSS 连接错误：", err)
    }

    log.Println("%#v\n", client)
    CreateBucket(constants.OSSVideoBucketName)
    CreateBucket(constants.OSSImgBucketName)
}
