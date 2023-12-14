package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func test(num int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("协程（%v）打印的第%v条数据\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
