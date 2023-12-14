package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	start := time.Now().Unix()

	//单向管道
	ch2 := make(chan<- int, 2) //只写管道
	ch2 <- 1
	ch2 <- 2
	//<- ch2
	fmt.Println(ch2)

	ch3 := make(<-chan int, 2) //只读管道
	//ch3 <- 1
	fmt.Println(ch3)

	end := time.Now().Unix()
	fmt.Printf("用时%vms\n", end-start)
}
