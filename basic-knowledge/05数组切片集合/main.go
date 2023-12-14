package main

import "fmt"

func main() {
	arrMain()
	sliceMain()
	mapMain()
}

// var array_name [quantity]Type
// array_name[index] = Value
func arrMain() {
	var exArray [3]int
	exArray[0] = 1
	fmt.Println(exArray[0])
	fmt.Println(len(exArray))
}

// var slice_name []Type
func sliceMain() {
	var resultSlice []int
	resultSlice = append(resultSlice, 2)
	resultSlice = append(resultSlice, 3)
	fmt.Println(resultSlice)
	fmt.Println("length: ", len(resultSlice))
	for index := range resultSlice {
		fmt.Println(resultSlice[index])
	}
}

// var map_name = make(map[key_type]value_type)
func mapMain() {
	var studentInfos = make(map[string]string)
	studentInfos["01"] = "Kevin"
	studentInfos["02"] = "Amy"
	fmt.Println(studentInfos)
	fmt.Println(studentInfos["01"])

	for num, name := range studentInfos {
		fmt.Println("学号：", num, " 姓名：", name)
	}
}
