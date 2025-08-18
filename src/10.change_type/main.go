package main

import (
	"fmt"
	"strconv"	//str convertの略
)

func main(){
	//変数名 = 型名(変換する変数名)

	var s string = "100"
	fmt.Println(s)
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)	//文字列型から数値型への変換.　_の意味は値を破棄という意味．この関数は2つの値を返すが，1つしか使用しないため，破棄する
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 100
	fmt.Println(i2)
	fmt.Printf("i2 = %T\n", i2)

	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)


}
