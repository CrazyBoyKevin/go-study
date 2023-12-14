package main

import (
	"fmt"
	"reflect"
)

func main() {
	d01()
	d02()
	typeTransfer()
}

func d01() {
	//exampleNumberA变量（整数型变量）声明和赋值
	var exampleNumberA int = 10
	//获取exampleNumberA的地址，并赋值给exampleNumberAPtr变量（exampleNumberAPtr的类型是指针类型）
	exampleNumberAPtr := &exampleNumberA
	//输出exampleNumberAPtr变量的值（将输出内存地址）
	fmt.Println(exampleNumberAPtr)
	//获取exampleNumberAPtr（指针变量）表示的实际数据值，并赋值给exampleNumberAPtrValue变量（整数型变量）
	exampleNumberAPtrValue := *exampleNumberAPtr
	//输出exampleNumberAPtrValue变量（整数型变量）的值
	fmt.Println(exampleNumberAPtrValue)
}

func d02() {
	//使用new()函数创建名为exampleNumberAPtr指针类型变量，表示int64型值
	exampleNumberAPtr := new(int64)
	//修改exampleNumberAPtr表示的实际数据值
	*exampleNumberAPtr = 100
	//获取exampleNumberAPtr表示的实际数据值
	fmt.Println(*exampleNumberAPtr)
}

func typeTransfer() {
	//声明float32型变量exampleFloat32并赋值
	var exampleFloat32 float32 = 150.25
	//将exampleFloat32转换为float64类型，并将结果赋值给exampleFloat64
	exampleFloat64 := float64(exampleFloat32)
	//输出exampleFloat64的类型和值
	fmt.Println(reflect.TypeOf(exampleFloat64), exampleFloat64)
	//将exampleFloat32转换为int32类型，exampleInt32
	exampleInt32 := int32(exampleFloat32)
	//输出exampleInt32的类型和值
	fmt.Println(reflect.TypeOf(exampleInt32), exampleInt32)
}
