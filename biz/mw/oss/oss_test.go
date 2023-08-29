package oss

import (
	"fmt"
	"testing"
	"tiktok_demo/pkg/constants"
)

func TestBucketExist(t *testing.T) {
    exists, err := client.IsBucketExist(constants.OSSVideoBucketName)
    if err != nil {
        fmt.Println(err)
        return
    }
    if exists {
        fmt.Println("%v found!\n", constants.OSSVideoBucketName)
    } else {
        fmt.Println("not found!")
    }
}

func TestBucketMake(t *testing.T) {
    exists, err := client.IsBucketExist(constants.OSSVideoBucketName)
    if err != nil {
        fmt.Println(err)
        return
    }
    if exists {
        fmt.Println("%v found!\n", constants.OSSVideoBucketName)
    } else {
        err = client.CreateBucket(constants.OSSVideoBucketName)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Successfully created bucket %v\n", constants.OSSVideoBucketName)
    }
}


