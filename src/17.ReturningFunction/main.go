package main

import "fmt"

// 関数を返す関数
// その場に合わせた柔軟な関数が作成できる

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// xと引数を足し合わせる関数を作成する関数
// func(int) int: これは整数を受け取って整数を返す型という意味
func addr(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main(){
	f := ReturnFunc()
	f()

	g := addr(10)	// gは10と引数を足し合わせる関数
	fmt.Println(g(5))

	double_g := addr(20)	//double_gは20と引数を足し合わせる関数
	fmt.Println(double_g(10))
}