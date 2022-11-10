package main

import (
	"fmt"
	"os"
)

type Movie struct {
	Name            string
	DirectorName    string
	WriterNames     []string
	CharactersNames []string
	Country         string
	Price           float64
	ShowTime        string
	MovieType       string
	Comments        []string
}

func main() {
	m := Movie{
		Name:            "崇文路历险记",
		DirectorName:    "高新波",
		WriterNames:     []string{"波比大王", "汇东车神", "豆腐魔", "黑下熄火"},
		CharactersNames: []string{"负数零", "黑大帅", "烈焰球", "大法师"},
		Country:         "中国大陆",
		Price:           0,
		ShowTime:        "2022年12月32日",
		MovieType:       "喜剧",
	}
	var i int
	var n float64
	var ThePrice float64
	var YourPrice float64
	n = 0
	ThePrice = 0
	for {
		fmt.Printf("请输入你的命令\n1：展示电影基本信息\n2:打分\n3:评论\n4：退出程序\n")
		fmt.Scanln(&i)
		if i == 1 {
			fmt.Println("电影名：", m.Name)
			fmt.Println("导演：", m.DirectorName)
			fmt.Println("主演：", m.WriterNames)
			fmt.Println("编剧：", m.CharactersNames)
			fmt.Println("地区：", m.Country)
			fmt.Println("类型：", m.MovieType)
			fmt.Println("上映时间：", m.ShowTime)
			if m.Price == 0 {
				fmt.Println("暂无人评分")
			} else {
				fmt.Printf("评分:%.1f\n", m.Price)
			}
			if m.Comments == nil {
				fmt.Println("暂无评论")
			} else {
				fmt.Println("评论", m.Comments)
			}
		} else if i == 2 {
			n++ //每次打分，人数加一
			fmt.Println("请打分：")
			fmt.Scanln(&YourPrice)
			// 多次打分求平均值
			ThePrice += YourPrice
			m.Price = ThePrice / n
		} else if i == 3 {
			var s string
			fmt.Println("请评论：")
			fmt.Scanln(&s)
			m.Comments = append(m.Comments, s)
		} else if i == 4 {
			fmt.Println("退出")
			os.Exit(0)
		}
		fmt.Print("\n")
	}
}
