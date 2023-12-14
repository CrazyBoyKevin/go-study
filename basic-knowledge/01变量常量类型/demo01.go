package main

// main包 每个程序有且只有一个

// 引入fmt包 包一旦引入必须被使用
import "fmt"

/*
 * 入口函数
 * 有且只有一个
 */
func main() {
	fmt.Println("Hello GO!")
	var num1 int = 100
	var num2 int = 900

	var (
		num3 int = 1
		num4 int = 2
	)
	const PI float64 = 3.14

	fmt.Println(num2 + num1)

	fmt.Println(num3 * num4)
}
