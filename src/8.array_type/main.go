package main

import "fmt"

func main(){
	var arr1 [3]int	//var 変数名　[要素数]要素の型名
	fmt.Println(arr1)	//[0 0 0]（整数型の初期値は0だから）
	fmt.Printf("%T\n", arr1)	//[3]intが型名となるため，Golangの配列は後から要素数を変更することができない

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)	//最後の要素は空文字

	arr3 := [3]int{1,2,3}	//暗黙的な宣言
	fmt.Println(arr3)

	arr4 := [...]string{"C","D"}	//[...]は与えられた要素数に自動で合わせてくれる
	fmt.Println(arr4)
	fmt.Printf("%T\n",arr4)	//型名は[2]stringになる（[...]stringではない）

	fmt.Println(arr2[0])	//値の取り出し（Pythonとかと変わらない）

	fmt.Println(arr3[3-1])	//要素数-1することで最後の値を取り出すことができる小技

	arr2[2] = "C"	//要素の更新

	fmt.Println(arr2)

	// var arr5 [4]int	
	// arr5 = arr1	//arr1とarr5の要素数が違うからエラーが出る
	// fmt.Println(arr5)

	fmt.Println(len(arr1))	//len関数は要素数を出力する（Pythonとかと一緒）
}