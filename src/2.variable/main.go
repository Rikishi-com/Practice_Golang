package main

import "fmt"

// i5 := 500	//暗黙的な定義は関数外では定義不可
var i5 int = 500	//明示的な定義は関数外でも定義可能

func outer(){
	var s4 string = "ueno"
	fmt.Println(s4)
}

func main(){
	var i int = 100	//var 変数名 型 = 値
	fmt.Println(i)

	var s string = "Hello World"
	fmt.Println(s)

	var t, f bool = true, false	//複数変数を一度に設定する（同じ型の場合）
	fmt.Println(t,f)

	var(		//複数変数を一度に設定する（異なる型の場合）
		i2 int = 200
		s2 string = "Hello Go"
	)
	fmt.Println(i2,s2)

	var i3 int	//int型は初期値として0と設定される
	var s3 string	//string型は初期値として空文字が設定される
	fmt.Println(i3,s3)	//0と表示される

	i3 = 300
	s3 = "Hello Lang"
	fmt.Println(i3,s3)

	i = 150	//値の更新
	fmt.Println(i)

	//暗黙的な定義（型宣言が必要ない，しかし型がないわけではない．型推論で勝手に形を決めてくれる）
	i4 := 400	//変数名 := 値
	fmt.Println(i4)

	i4 = 450	//値の更新は暗黙的な定義でも同様の方法
	fmt.Println(i4)

	// i4 = "Hello"  //型推論によりi4はintとされているから，エラーが出る
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()	//outerは別関数

	// fmt.Println(s4)	//変数は関数内でのみ有効なのでエラーが出る

	// var s5 string = "yuki"	//Golangでは定義された変数は必ず使用する必要がある


}