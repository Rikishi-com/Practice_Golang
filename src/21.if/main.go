package main

import "fmt"

func main(){
	a := 0

	if a == 2 {		// 波線はswitchも使えるというだけ
		fmt.Println("No")
	} else if a == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("I Don't Know")
	}

	// 簡易if文
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 1000

	if x := 2; true {	// if文内で定義した変数はifブロック内のみで有効
		fmt.Println(x)
	}
	fmt.Println(x)

	// fmt.Println(b)	// bはif文外で定義されていないからエラーになる
	
}
