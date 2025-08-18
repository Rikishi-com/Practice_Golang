package main

import "fmt"

func main(){
	var fl64 float64 = 1.1
	fmt.Println(fl64)

	fl := 2.2	//浮動小数点型の暗黙的な定義の場合，自動でfloat64になる（環境依存ではない）

	fmt.Println(fl)

	fmt.Printf("%T\n%T\n", fl,fl64)

	fmt.Println(fl + fl64)	//同じ型なので演算が可能（しかし丸め誤差が出る．これはどの言語でも一緒なので桁数制限すればOK）

	// var ueno float = 1.5	//floatという型はない

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)	//+inf = 正の無限大（infinity）

	ninf := - 1.0 / zero
	fmt.Println(ninf)	//-inf = 負の無限大(infinity)

	nan := zero / zero
	fmt.Println(nan)	//NaN = Not a Number = 定義できない

}