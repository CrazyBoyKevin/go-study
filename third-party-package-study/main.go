package main

//go get 全局安装
//go mod download
//go mod vendor	将依赖复制到项目的vendor下

import (
	"github.com/tidwall/gjson"
	"github.com/shopspring/decimal"
	"fmt"
)


const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())

	quantity := decimal.NewFromInt(3)
	fmt.Println(quantity)
}