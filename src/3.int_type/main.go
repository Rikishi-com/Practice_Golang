package main

import "fmt"

func main(){
	var i int = 100	//intの桁数は環境依存（使用環境が64bitで動いているならint64で動作する）

	var i2 int64 = 100	//明示的にすることも可能（int8,16,32,64などがある）

	fmt.Println(i + 50)

	// fmt.Println(i + i2)	
	//同じ整数型でもビット数が異なると演算はできない（基本的なPCは64bitだが，環境依存の64bitと明示的な64bitでは全く異なる型と判定される）

	fmt.Printf("%T\n", i)	//%Tは書式指定子．Tは型を表す

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))	//型変換は　型名(変数名)

}