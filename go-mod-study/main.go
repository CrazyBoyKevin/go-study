package main

import (
	"go-mod-study/calc"
	"fmt"
	"go-mod-study/tools"
)

func init() {
	fmt.Println("我先init")
}

func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)
	result := calc.Sub(10, 2)
	fmt.Println(result)
	fmt.Println(calc.Age)

	//调用tools里面的方法
	b := tools.Mul(2,3)
	fmt.Println(b)
}