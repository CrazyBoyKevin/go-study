package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var rwMutex = sync.RWMutex{}

// goroutine 读写互斥锁
// 读的时候 完全互斥
func write() {
	rwMutex.Lock()
	fmt.Println("write")
	time.Sleep(time.Second * 2)
	rwMutex.Unlock()
	wg.Done()
}

// 写的时候并行操作
func read() {
	rwMutex.RLock()
	fmt.Println("read")
	time.Sleep(time.Second * 2)
	rwMutex.RUnlock()
	wg.Done()
}

func main() {
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}
