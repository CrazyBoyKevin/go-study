package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 让主线程等待所有协程执行完毕后退出
var wg = sync.WaitGroup{}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test hello")
		time.Sleep(time.Millisecond * 100)
	}
	//通知wg我执行结束了 协程计数器-1
	wg.Done()
}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1 hello")
		time.Sleep(time.Millisecond * 100)
	}
	//通知wg我执行结束了 协程计数器-1
	wg.Done()
}

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Printf("我的CPU个数：%d\n", cpuNum)
	//设置程序占用的cpu个数，默认是全部
	runtime.GOMAXPROCS(cpuNum - 2)
	//go 表示开启一个 协程
	//添加一个协程计数
	wg.Add(1)
	go test()
	//添加一个协程计数
	wg.Add(1)
	go test1()
	//主进程执行完毕，无论协程是否运行完成，程序都会退出，不等待协程执行完毕
	for i := 0; i < 10; i++ {
		fmt.Println("main hello")
		time.Sleep(time.Millisecond * 50)
	}
	//执行完毕后退出
	wg.Wait()
}
