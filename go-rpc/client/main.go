package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsRes struct {
	Success bool
	Msg     string
}

type GetGoodsReq struct {
	Id int
}

type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func main() {
	//1.用rpc.Dial和rpc微服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8801")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//2.当客户端推出的时候 关闭连接
	defer client.Close()
	//3.远程调用函数
	var reply GetGoodsRes
	err2 := client.Call("goods.GetGoods", GetGoodsReq{
		Id: 1,
	}, &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//4.获取微服务返回的数据
	fmt.Println(reply)
}
