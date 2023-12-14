package main

import (
	"sync"
	"time"
)

var count = 0
var wg = sync.WaitGroup{}
var mutex sync.Mutex

var m1 = make(map[int]int, 0)

// goroutine 互斥锁
func test(num int) {
	mutex.Lock()
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m1[num] = sum
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for i := 1; i < 40; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
