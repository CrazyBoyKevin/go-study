package main

import "fmt"

type Animal struct {
	Name   string
	Age    int
	Gender string
}

type Bird struct {
	WingColor    string
	CommonAnimal Animal
}

func NewBird(name string, age int, gender string, wingColor string) *Bird {
	return &Bird{
		WingColor: wingColor,
		CommonAnimal: Animal{
			Name:   name,
			Age:    age,
			Gender: gender,
		},
	}
}

func (b *Bird) Fly() {
	fmt.Println("我起飞啦！")
}

func main() {
	bird := *NewBird("小鸟", 1, "公", "绿色")
	fmt.Println(bird)
	bird.Fly()
}
