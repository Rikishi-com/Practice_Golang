package main

import (
	"fmt"
	"math"
)

// 無名関数
// 初期化の即時実行・コールバック・goroutine内の
// 一時処理・外部変数を保持するクロージャなど、
// 局所的な処理をその場で手短にまとめたいときに使う

func main(){
	f := func(x, y int) int {
		return x + y
	}

	i := f(1,2)
	fmt.Println(i)

	//　そのまま値も入れてしまう省略記法
	i2 := func(p,q float64) float64 {
		return math.Pow(p,q)
	}(2,4)

	fmt.Println(i2)
}
