package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//创建远程调用的函数，函数一般是放在结构体里面

type Goods struct {
}

// AddGoodsReq AddGoods参数对应的结构体
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

func (good *Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	//1.查询执行
	fmt.Println(req)
	*res = GetGoodsRes{
		Id:      1,
		Title:   "title",
		Price:   99.99,
		Content: "content",
	}
	//返回结果
	return nil
}

func (good *Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	//1.增加执行
	fmt.Println(req)
	*res = AddGoodsRes{
		Success: true,
		Msg:     "success",
	}
	//返回结果
	return nil
}

func main() {
	//1.注册RPC服务
	err := rpc.RegisterName("goods", new(Goods))
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8801")
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
		//rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //使用json格式传输
	}
}
