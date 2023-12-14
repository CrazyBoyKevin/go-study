package main

import "fmt"

/*
*

	func function_name([params_list])([return_values_list]) {
		// 函数体
	}
*/
func main() {
	//多个defer存在时，它们的顺序是反向的
	defer fmt.Print("素数")
	defer fmt.Print("查找")
	result := stringLoop("字符串回环测试")
	fmt.Println(result)
}

func stringLoop(content string) (string) {
	return content
}
ti