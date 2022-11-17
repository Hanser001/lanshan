package main

import (
	"fmt"
	"time"
)

var Wallet = 0
var ch1 = make(chan struct{}, 1)

func vPaopao50() {
	ch1 <- struct{}{}
	Wallet += 50
	<-ch1
}

func main() {
	for i := 0; i < 10000; i++ {
		go vPaopao50()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("泡泡现在有", Wallet, "元")
}
