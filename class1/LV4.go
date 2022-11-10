package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// 作弊模式
	//fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	// 通过一个 for 循环实现一直猜数，直到猜中
	for {
		var guessGroup = make([]int, 5)
		//用一个切片装多个数字，提高猜中几率
		fmt.Println("输入五个数字")
		for i := 0; i < 5; i++ {
			_, err := fmt.Scan(&guessGroup[i])
			if err != nil {
				fmt.Println("Invalid input. Please enter an integer value")
				continue
			}
		}
		for i := 0; i < 5; i++ {
			fmt.Println("You guess is", guessGroup[i])
			if guessGroup[i] > secretNumber {
				fmt.Println("Your guess is bigger than the secret number. Please try again")
			} else if guessGroup[i] < secretNumber {
				fmt.Println("Your guess is smaller than the secret number. Please try again")
			} else {
				fmt.Println("Correct, you Legend!")
				break
			}
		}
	}
}
