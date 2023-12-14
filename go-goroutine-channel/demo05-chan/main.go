package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	//关闭管道，否则for range的时候会deadlock,for i的时候没事，建议关闭
	close(ch1)
	//管道没有key
	for val := range ch1 {
		fmt.Println(val)
	}
}
