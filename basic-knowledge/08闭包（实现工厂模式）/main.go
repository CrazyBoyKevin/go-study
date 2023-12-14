package main

import "fmt"

// 游戏角色创建器，传入角色名、职业和性别
func createPlayer(name string, career string, gender string) func() (string, string, string, int, int) {
	var hp = 0
	var mp = 0
	if career == "战士" {
		hp = 150
		mp = 100
	} else if career == "法师" {
		hp = 100
		mp = 200
	}
	// 闭包
	return func() (string, string, string, int, int) {
		return name, career, gender, hp, mp
	}
}

func main() {
	//工厂模式
	playerA := createPlayer("狂斩天下", "战士", "男")
	nameA, careerA, genderA, hpA, mpA := playerA()
	fmt.Println(genderA, careerA, nameA, hpA, mpA)

	playerB := createPlayer("温玉琳琅", "法师", "女")
	nameB, careerB, genderB, hpB, mpB := playerB()
	fmt.Println(genderB, careerB, nameB, hpB, mpB)

	//闭包
	//发生战斗，B消耗15点魔法值
	mpB -= 15
	//发生战斗，A失去20点血量值
	hpA -= 20
	fmt.Println(nameB, hpB, mpB)
	fmt.Println(nameA, hpA, mpA)

}
