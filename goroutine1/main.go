package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	
	fmt.Println("普通に関数を呼び出す")
	serialno()
	serialno()

	fmt.Println("ゴルーチンとして呼び出す")
	go serialno()
	go serialno()
	go serialno()

	//ゴルーチン呼び出し後、sleepする
	time.Sleep(1 * time.Second)
		
	fmt.Println("main end...")
}

func serialno() {
	for i := 0; i < 5; i ++ {
		fmt.Println(i)
		//1秒間Sleepする
		time.Sleep(1 * time.Second)
	}
}
