package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Hello 定义一个远程调用的方法
type Hello struct {
}

/*
SayHello
1.方法只能有两个可序列化的参数，其中第二个参数是指针类型
--req 表示客户端传过来的数据
--res 表示给客户端返回的数据
2.方法要返回一个error类型，同时必须是公开的方法
3.req和res的类型不能为 channel\func\complex(不能序列化)
*/
func (this *Hello) SayHello(req string, res *string) error {
	fmt.Println(req)
	*res = "hello! " + req
	return nil
}

func main() {
	//1.注册RPC服务
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println(err)
		return
	}
	//3.应用推出的时候关闭监听
	defer listener.Close()

	//4.建立连接
	for {
		fmt.Println("开始建立连接")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//5.绑定服务
		rpc.ServeConn(conn)
	}

}
