package main

import (
	"fmt"
	"math"
)

func main() {
	// 加算演算子
	fmt.Println(1 + 2)

	// 減算演算子
	fmt.Println(5 - 1)

	// 乗算演算子
	fmt.Println(5 * 3)

	// 除算演算子
	fmt.Println(60 / 3)

	// 余り演算子
	fmt.Println(9 % 4)

	// 文字列連結演算子
	fmt.Println("ueno" + "yuki")

	n := 0
	// 加算代入演算子
	n += 4 //n = n + 4

	fmt.Println(n)

	// インクリメント演算子
	n++ //n = n + 1
	fmt.Println(n)

	// デクリメント演算子
	n-- //n = n - 1
	fmt.Println(n)

	s := "ABC"
	// 文字列連結代入演算子
	s += "DEF" //s = ABC + DEF

	fmt.Println(s)

	// べき乗（math.Pow を使用）
	// math パッケージをインポートする必要がある
	// math.Pow は float64 を返すので、整数計算に使う場合は注意
	// べき乗
	fmt.Println(math.Pow(2, 3))  // 2の3乗 = 8
	fmt.Println(math.Pow(9, 0.5)) // 平方根 = 3
}
