package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 100
	lock           = sync.Mutex{}
	moneyDeposited = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		fmt.Println("Stingy sees balance of ", money)
		moneyDeposited.Broadcast()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy(whoAmI int) {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			moneyDeposited.Wait()
		}
		money -= 20
		fmt.Printf("%d - Spendy sees balance of %d \n", whoAmI, money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	//go stingy()
	//go spendy(0)
	for i := 0; i <= 2; i++ {
		go stingy()
	}
	for i := 0; i <= 5; i++ {
		go spendy(i)
	}
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
