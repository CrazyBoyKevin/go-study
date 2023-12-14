package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// 写数据
func fn1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("写入%v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
	wg.Done()
}

// 读数据
func fn2(ch chan int) {
	for val := range ch {
		fmt.Printf("读取%v\n", val)
	}
	wg.Done()
}

func main() {
	ch1 := make(chan int, 10)
	wg.Add(1)
	go fn1(ch1)
	wg.Add(1)
	go fn2(ch1)
	wg.Wait()

	fmt.Println("exit")
}
