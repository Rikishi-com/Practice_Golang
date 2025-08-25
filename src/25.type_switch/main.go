package main

import "fmt"

// 型スイッチ

func anything(a interface{}){
	fmt.Println(a)
}

func main(){
	anything("aaa")
	anything(1)

	var x interface{} = 3	// interface型は計算できない
	i := x.(int)	// 型アサーション（型復元）
	fmt.Println(i + 2)	// 整数型に復元することで計算できるようになった

	// f := x.(float64)	//エラーになってできない

	f , isFloat64 := x.(float64)	// できない場合には2つ目にFalseが返ってくる
	fmt.Println(f,isFloat64)

	// interface型の便利な型判定
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}

	// interface型の型アサーション
	switch v := x.(type){
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	case bool:
		fmt.Println(v, "bool")
	default:
		fmt.Println("other type")
	}

}