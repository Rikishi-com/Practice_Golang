package main

import "fmt"

func main(){
	var x interface{}	//すべての型と互換性を持つ特殊な型
	fmt.Println(x)	//nilが返る．nilはnoneと同じ意味

	x = 1
	fmt.Println(x)

	x = "ueno"
	fmt.Println(x)

	x = [3]int{1,2,3}	
	fmt.Println(x)

	// x = 2	//型特有の演算ができない
	// fmt.Println(x + 2)
}
