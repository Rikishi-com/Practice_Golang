package main

import "fmt"

// 式スイッチ

func main(){
	n := 1

	switch n {
	case 1,2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch n := 3; n {	// ifと同様にこの関数内でのみ使用可（変数の局所性を高められる）
	case 1,2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	// x := 6

	// switch {
	// case n > 0 && x < 4:	// 列挙型と演算子型を同時に使用できない
	// 	fmt.Println(" 0 < x < 4")
	// case 3,4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }
}