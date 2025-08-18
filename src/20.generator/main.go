package main

import "fmt"

// ジェネレーター
// 何らかの規則に従って連続した値を返す関数
// クロージャーの応用

func intergers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// xから始まるインクリメント
func argment_intergers(x int) func() int {
	i := x - 1
	return func() int {
		i++
		return i
	}
}

func main(){
	ints := intergers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	fmt.Println("---")

	arg_ints := argment_intergers(3)
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())

	fmt.Println("---")

	// 関数を新しく宣言したら新しくカウントする
	second_arg_ints := argment_intergers(3)
	fmt.Println(second_arg_ints())
	fmt.Println(second_arg_ints())
	fmt.Println(second_arg_ints())

	fmt.Println("---")

	// ここで使っても引き継がれる
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())



}