// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"tiktok_demo/biz/dal"
)

func Init() {
    dal.Init()
}

func main() {
	h := server.Default(
        server.WithStreamBody(true),
        server.WithHostPorts("0.0.0.0:18005"),
    )
	register(h)
	h.Spin()
}