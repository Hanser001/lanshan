package main

import "fmt"

func main() {
	//n是数组长度，i和r是区间
	var n int
	var i int
	var r int
	var m int
	fmt.Scanln(&n, &i, &r)
	s := make([]int, n)
	//把元素装进数组
	for m := 0; m < n; m++ {
		fmt.Scan(&s[m])
	}

	//外层循坏表示要进行多少次（r-l）次排序，内层循环是每次将最小的数放到前面
	for j := r; j >= i; j-- {
		for i := i; i <= r; i++ {
			if s[i] > s[i+1] {
				m = s[i+1]
				s[i+1] = s[i]
				s[i] = m
			}
		}
	}
	fmt.Println(s)
}
