package main

import (
	"fmt"
	"time"
)

var (
	money          = 100
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		money += 10
		fmt.Println("Stingy sees balance of ", money)
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		money -= 20
		fmt.Println("Spendy sees balance of ", money)
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
