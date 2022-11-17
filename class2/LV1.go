package main

import (
	"fmt"
	"time"
)

// 定义一个函数类型
type functype func(string)

func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() {
	fmt.Println("输了啦，都是你害的.")
}

// 函数类型作为参数传入
func take(type1 functype) {
	好康的 := 欢迎来我家玩()
	type1(好康的)
}

func main() {
	//使用一个匿名函数
	go take(func(好康的 string) {
		fmt.Println(好康的)
	})
	打电动()
	//保证take函数在main函数结束前执行完毕
	time.Sleep(6 * time.Second)
}
