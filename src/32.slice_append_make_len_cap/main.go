package main

import "fmt"

func main(){
	sl := []int{100,200}
	fmt.Println(sl)

	// append 要素を追加
	sl = append(sl,300)
	fmt.Println(sl)

	// 複数OK
	sl = append(sl,400,500)
	fmt.Println(sl)

	// make 初期値で作成
	sl2 := make([]int,5)
	fmt.Println(sl2)

	// sl3 := make([]int,5 ,10)		// 第二引数は容量

	// len 要素数を返す

	// cap 容量を返す

}