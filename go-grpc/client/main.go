package main

import (
	"client/proto/greeter"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//1.建立连接; grpc.WithTransportCredentials(insecure.NewCredentials()) --> 连接的安全性
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	//2.注册客户端
	client := greeter.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "I'm Client",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", res)
	fmt.Println(res.Message)
}
