package main

import "fmt"

/**
//定义结构体的标准格式
type StructName struct {
	//属性字段
}
*/

type Player struct {
	// 角色名
	name string
	// 职业
	career string
	// 性别
	gender string
}

func anonymousStruct() {
	wishingTree := struct {
		height   float64
		width    float64
		treeType string
	}{22.5 * 100, 50, "banyan"}
	fmt.Println(wishingTree)
}

type Dog struct {
	Breed  string
	Age    int
	Weight float64
	Gender string
}

func NewDog(breed string, age int, weight float64, gender string) *Dog {
	return &Dog{
		Breed:  breed,
		Age:    age,
		Weight: weight,
		Gender: gender,
	}
}

func main() {
	player := Player{"Kevin", "programmer", "man"}
	player.name = "KEVIN"
	fmt.Println(player.name)
	anonymousStruct()

	fatShibaInu := NewDog("Shiba Inu", 2, 12.0, "公")
	weakShibaInu := NewDog("Shiba Inu", 2, 7.0, "公")
	fmt.Println(fatShibaInu)
	fmt.Println(weakShibaInu)
}

/**
//方法
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    方法体
}
*/
