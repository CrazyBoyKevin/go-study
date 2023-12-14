package main

import (
	"fmt"
	"time"
)

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func test2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {
	go test1()
	go test2()
	time.Sleep(time.Second)
}
