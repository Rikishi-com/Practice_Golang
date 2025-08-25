package main

import "fmt"

func main(){
	sl := []int{100,200}
	sl2 := sl	// 参照型はアドレスコピー
	fmt.Println(sl2)

	var i int = 100	// 基本型はアドレスコピーではない
	i2 := i
	i2 = 200

	fmt.Println(i,i2)

	ssl := []int{1,2,3}
	ssl2 := make([]int,3)	// 空の配列を用意しておく

	n := copy(ssl2,ssl)	// 第一引数：コピー先，第二引数：コピー元　n:コピーした要素数

	fmt.Println(n,ssl2)

}
