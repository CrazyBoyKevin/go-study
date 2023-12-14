package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// 写数据
func putNum(ch chan int) {
	for i := 0; i < 100000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

// 读数据
func primeNum(intChan chan int, primeChan chan int, exitChan chan int) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num //num是素数
		}
	}
	//需要关闭，如果关闭了那就不能往chan里面写数据了
	exitChan <- 1
	wg.Done()
}

func printPrime(primeChan chan int) {
	for val := range primeChan {
		fmt.Println(val)
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	intChane := make(chan int, 100000)
	primeChan := make(chan int, 100000)
	exitChan := make(chan int, 16) //标识primeChan close
	//存放
	wg.Add(1)
	go putNum(intChane)
	for i := 0; i < 16; i++ {
		//统计
		wg.Add(1)
		go primeNum(intChane, primeChan, exitChan)
	}
	//打印
	wg.Add(1)
	go printPrime(primeChan)

	//判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		//关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("用时%vms\n", end-start)
}
