package main

import (
	"fmt"
)

// 可変長引数

func Sum(s ...int) int {	// 引数の数何個でもOK
	n := 0
	for _,v := range s {
		n += v
	}
	return n
}

func main(){
	fmt.Println(Sum(1,2,3,4,5,6,7,8,9))
	fmt.Println(Sum())
}