package main

import "fmt"

// channel
// 複数のゴルーチン間でデータの受け渡しをするために設計されたデータ構造
// 宣言，操作

func main(){
	var ch1 chan int

	// 受信専用
	// var ch2 <- chan int

	// 送信専用
	// var ch3 chan <- int

	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))	// capは容量計算
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	ch3 <- 2
	ch3 <- 3

	for i := 0; i < 3; i++ {
		fmt.Println(<- ch3)	// 1,2,3の順で出てくる
	}
}