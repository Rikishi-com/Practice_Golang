package main

import "fmt"

func main() {

	// 等値演算子（等しい）
	fmt.Println(1 == 1) //true

	// 等値演算子（等しくない場合）
	fmt.Println(1 == 2) //false

	// 不等値演算子（等しくない）
	fmt.Println(1 != 2) //true

	// 不等値演算子（等しい場合）
	fmt.Println(1 != 1) //false

	// 小なり演算子
	fmt.Println(3 < 5) //true

	// 大なり演算子
	fmt.Println(5 > 3) //true

	// 小なりイコール演算子
	fmt.Println(4 <= 5) //true

	// 小なりイコール演算子（等しい場合）
	fmt.Println(5 <= 5) //true

	// 大なりイコール演算子
	fmt.Println(5 >= 4) //true

	// 大なりイコール演算子（等しい場合）
	fmt.Println(5 >= 5) //true

	// 文字列の比較
	fmt.Println("abc" == "abc") //true
	fmt.Println("abc" != "def") //true
}
