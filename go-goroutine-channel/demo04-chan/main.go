package main

import "fmt"

func main() {
	//1创建管道
	ch := make(chan int, 3)
	//2给管道里面存储内容
	ch <- 10
	ch <- 21
	ch <- 32
	fmt.Printf("值:%v,容量:%v,长度%v\n", ch, cap(ch), len(ch))
	//3获取管道里面的内容
	a := <-ch
	fmt.Println(a)
	b := <-ch
	fmt.Println(b)
	c := <-ch
	fmt.Println(c)
	//管道里面的数据遵循 先入先出
	fmt.Println("---------------------------------")

	//管道属于引用数据类型
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64
	ch2 := ch1
	ch2 <- 25
	<-ch1
	<-ch1
	<-ch1
	fmt.Println(<-ch1)

	//管道的阻塞，管道放满了放不下了 就是阻塞 fatal error: all goroutines are asleep - deadlock!
	ch6 := make(chan int, 1)
	ch6 <- 34
	//ch6 <- 64

	//管道没有数据了，还在取，也是阻塞 fatal error: all goroutines are asleep - deadlock!
	ch7 := make(chan string, 3)
	ch7 <- "num1"
	ch7 <- "num2"
	m1 := <-ch7
	m2 := <-ch7
	//m3 := <-ch7
	//fmt.Println(m1, m2, m3) //为了不报错 注释了
	fmt.Println(m1, m2)
}
