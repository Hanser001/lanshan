package main

import (
	"fmt"
	"time"
)

// 3个chan用来控制开火的三个流程顺序
var (
	chan1 = make(chan struct{}, 1)
	chan2 = make(chan struct{}, 1)
	chan3 = make(chan struct{}, 1)
)

// 士兵数是常量
const (
	loadMan = 10
	aimMan  = 5
	fireMan = 3
)

// 自定义函数内部用for循环让每个士兵能重复工作
func loadRounds() {
	for {
		<-chan1
		//装弹花一秒
		time.Sleep(time.Second)
		fmt.Print("装弹->")
		chan2 <- struct{}{}
	}
}

func takeAim() {
	for {
		//保证先装弹后瞄准
		<-chan2
		//瞄准花两秒
		time.Sleep(2 * time.Second)
		fmt.Print("瞄准->")
		chan3 <- struct{}{}
	}
}

func fire() {
	for {
		//保证先瞄准，后开火
		<-chan3
		//花一秒发射
		time.Sleep(time.Second)
		fmt.Print("发射!\n")
		chan1 <- struct{}{}
	}
}

func main() {
	// 开始装弹
	chan1 <- struct{}{}
	for i := 0; i < loadMan; i++ {
		go loadRounds()
	}
	for i := 0; i < aimMan; i++ {
		go takeAim()
	}
	for i := 0; i < fireMan; i++ {
		go fire()
	}
	//堵塞（不懂怎么用监听键盘的包QAQ）
	for {
		fmt.Print("")
	}
}
