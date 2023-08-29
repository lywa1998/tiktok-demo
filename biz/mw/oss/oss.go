package oss

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/url"
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
    fileObj, _ := file.Open()
    defer fileObj.Close()
    err = client.PutObject(bucketName + file.Filename, fileObj)
    return err
}

// GetObjURL get the original link of the file in oss
func GetObjURL(objectName string, expTime int64) (u *url.URL, err error) {
    u, err = client.signURL(objectName, oss.HTTPGet, expTime)
    return u, err
}

func PutToBucketByBuf(dst string, buf *bytes.Buffer) (err error) {
    err = oss.PutObject(dst, buf)
    return
}

func PutToBucketByFilePath(dst, src string) (err error) {
    err = oss.PutObjectFromFile(dst, src)
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
