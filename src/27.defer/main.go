package main

import (
	"fmt"
	"os"
)

// defer
// 関数がreturnする直前に必ず実行される処理を登録するためのキーワード

func TestDefer(){
	defer fmt.Println("END")	// これが最後に実行される
	fmt.Println("START")
}

func RunDefer(){
	defer fmt.Println("1")	//3
	defer fmt.Println("2")	//2
	defer fmt.Println("3")	//1
}

func main(){
	TestDefer()	// 1

	defer func() {	// last
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()	// 2


	file , err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()	// 必ず最後に閉じる必要があるから

	file.Write([]byte("Hello"))
}