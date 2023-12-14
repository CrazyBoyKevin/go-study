package main

import (
	"fmt"
	"go-protobuf/proto/userService"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := &userService.Userinfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"吃饭", "睡觉", "写代码"},
	}
	fmt.Println(u)
	fmt.Println(u.GetHobby())
	//Protobuf的序列化
	marshal, _ := proto.Marshal(u)
	fmt.Println(marshal)
	//Protobuf的反序列化
	user := userService.Userinfo{}
	proto.Unmarshal(marshal, &user)
	fmt.Printf("%#v\n", user)
	fmt.Println(user.GetType())
}
