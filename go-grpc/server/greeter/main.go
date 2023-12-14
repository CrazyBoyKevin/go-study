package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greeter/proto/greeter"
	"net"
)

// 1.定义远程调用的结构体和方法，我们这个结构图需要实现 GreeterServer的接口

type Hello struct {
}

func (this Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "你好" + req.Name,
	}, nil
}

func main() {
	//1.初始化一个grpc的对象
	server := grpc.NewServer()
	//2.注册服务
	greeter.RegisterGreeterServer(server, &Hello{})
	//3.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	//4.启动服务
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
