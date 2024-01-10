package handler

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	hello "microHello/service/hello/proto"
)

func HandleHello(c *gin.Context) {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	service := micro.NewService(
		micro.Registry(etcdReg), //设置注册中心
	)
	service.Init()

	client := hello.NewHelloService("hello", service.Client())
	// 调用服务
	//rsp, err := client.Call(c, &hello.CallRequest{
	//	Name: c.Query("key"),
	//})

	rsp, err := client.HandleHello(c, &hello.HelloRequest{UserName: c.Query("userName")})

	if err != nil {
		c.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": rsp.Msg})
}
