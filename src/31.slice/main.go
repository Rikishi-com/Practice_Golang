package main

import "fmt"

// スライス
// 宣言，操作
// 配列に似ている
func main(){
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100,200}
	fmt.Println(sl2)

	sl3 := []string{"A","B"}
	fmt.Println(sl3)

	sl4 := make([]int,5)	// [0 0 0 0 0]
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1,2,3,4,5}
	fmt.Println(sl5[2])	// 3

	fmt.Println(sl5[2:4])	//[3,4]

	fmt.Println(sl5[:2])	// :の手前までという意味(インデックス番号2の手前までなので[1 2])	// 未満

	fmt.Println(sl5[3:])	// 以上
}