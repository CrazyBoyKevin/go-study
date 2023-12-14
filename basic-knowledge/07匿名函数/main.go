package main

import (
	"fmt"
	"time"
)

/*
	func ([params_list]) (return_values_list) {
		// 函数体
	}
*/
var percent = 0
var keepChecking = true

func download(setStatus func()) {
	for {
		time.Sleep(1 * time.Second)
		percent++
		if percent == 50 {
			setStatus()
			break
		}
	}
}

func main() {
	setStatus := func() {
		keepChecking = false
	}

	go download(setStatus)

	for {
		if keepChecking {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("当前进度：", percent)
		} else {
			break
		}
	}

	// 定义匿名函数
	//func(text string) {
	//	fmt.Println(text)
	//}("定义时就调用")
	// 定义匿名函数
	//exampleVal := func(text string) {
	//	fmt.Println(text)
	//}
	//exampleVal("通过变量调用匿名函数")
}
