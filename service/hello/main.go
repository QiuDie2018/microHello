package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"hello/handler"
	pb "hello/proto"
)

var (
	EtcdAddr = "127.0.0.1:2379"
	service  = "hello"
	version  = "latest"
)

func main() {
	// Create service
	etcdReg := etcd.NewRegistry(
		registry.Addrs(EtcdAddr),
	)
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(etcdReg),
	)

	// Register
	if err := pb.RegisterHelloHandler(srv.Server(), new(handler.Hello)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
