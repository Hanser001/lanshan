package main

import (
	"fmt"
)

// 接口部分
type Sayer interface {
	Say()
}

type Attacker interface {
	Attack()
}

type Skiller interface {
	UseSkill()
}

type Proper interface {
	UseProp(prop string, agkAttend int)
}

// 结构体（英雄）部分
type Kaisa struct {
	HP  int
	MP  int
	Agk int
}

type Vayne struct {
	HP  int
	MP  int
	Agk int
}

// 方法部分
func (k Kaisa) Say() {
	fmt.Println("卡莎:我回来，是为了那些回不来的人")
}

func (k Kaisa) Attack() {
	fmt.Printf("卡莎的平A造成了%d点伤害\n", k.Agk)
}

func (k Kaisa) UseSkill() {
	fmt.Printf("卡莎使用了艾卡西亚暴雨，对敌人造成了%d点伤害\nMP减少了10点\n", k.Agk*2)
	k.MP -= 10
}

func (k *Kaisa) UseProp(prop string, agkAttend int) {
	fmt.Printf("卡莎使用%s，攻击力增加了%d点\n", prop, agkAttend)
	k.Agk += agkAttend
}

func (v Vayne) Say() {
	fmt.Println("薇恩:木已成舟")
}

func (v Vayne) Attack() {
	fmt.Printf("薇恩的平A造成了%d点伤害\n", v.Agk)
}

func (v *Vayne) UseProp(prop string, agkAttend int) {
	fmt.Printf("薇恩使用%s，攻击力增加了%d点\n", prop, agkAttend)
	v.Agk += agkAttend
}

func (v *Vayne) UseSkill() {
	fmt.Printf("薇恩使用终极时刻，攻击力增加50点\nMP减少30点\n")
	v.Agk += 50
	v.MP -= 30
}

// 实现接口部分
func showHero(myHero Sayer) {
	myHero.Say()
}

func heroAttack(myHero Attacker) {
	myHero.Attack()
}

func heroUseSkill(myHero Skiller) {
	myHero.UseSkill()
}

func heroUseProp(myHero Proper) {
	myHero.UseProp("愤怒合剂", 50)
}

// 二级菜单
func menu2() {
	fmt.Println("欢迎来到召唤师峡谷")
	fmt.Printf("请输入下一步指令\n1.平A\n2.使用技能\n3.使用愤怒合剂\n4.退出游戏,回到客户端\n")
}

func main() {
	hero1 := Kaisa{
		HP:  310,
		MP:  200,
		Agk: 79,
	}
	hero2 := Vayne{
		HP:  320,
		MP:  200,
		Agk: 75,
	}
	var option1 int
	var option2 int
	var option3 int
	heroGroup := []string{"Kaisa", "Vayne"}

Game:
	for {
		fmt.Printf("请输入你的命令\n1.选择你的英雄\n2.获取英雄列表\n3.退出程序\n")
		fmt.Scanln(&option1)
		switch option1 {
		case 1:
			fmt.Println("英雄列表:")
			for i := 0; i < len(heroGroup); i++ {
				fmt.Println(heroGroup[i])
			}
			fmt.Scanln(&option2)
			if option2 == 1 {
				showHero(hero1)
				menu2()
				for {
					fmt.Scanln(&option3)
					if option3 == 1 {
						heroAttack(hero1)
					} else if option3 == 2 {
						heroUseSkill(hero1)
					} else if option3 == 3 {
						heroUseProp(&hero1)
					} else if option3 == 4 {
						fmt.Println("退出本局游戏，回到客户端")
						goto Game
					}
				}
			} else if option2 == 2 {
				showHero(hero2)
				menu2()
				for {
					fmt.Scanln(&option3)
					if option3 == 1 {
						heroAttack(hero2)
					} else if option3 == 2 {
						heroUseSkill(&hero2)
					} else if option3 == 3 {
						heroUseProp(&hero2)
					} else if option3 == 4 {
						fmt.Println("退出本局游戏，回到客户端")
						goto Game
					}
				}
			} else {
				fmt.Println("无效输入,返回")
			}
		case 2:
			for i := 0; i < len(heroGroup); i++ {
				fmt.Println(heroGroup[i])
			}
		case 3:
			fmt.Println("退出游戏")
			break Game
		}
	}
}
