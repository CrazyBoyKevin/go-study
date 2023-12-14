package main

import "fmt"

/*
for init; condition; post {
    //循环体代码块
}
*/

func main() {
	n := 7
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
