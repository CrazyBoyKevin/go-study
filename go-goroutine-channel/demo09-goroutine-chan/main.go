package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//使用select多路复用获取chan的时候 不用关闭chan
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intChan %v\n", v)
			time.Sleep(time.Millisecond * 100)
		case v := <-stringChan:
			fmt.Printf("stringChan %v\n", v)
			time.Sleep(time.Millisecond * 100)
		default:
			fmt.Println("exit")
			return
		}
	}
}
