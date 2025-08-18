package main

import "fmt"

// 定数
// 頭文字を大文字にすることで，他のパッケージからも呼び出せる
// このパッケージ内だけなら，小文字でも良い
// 再定義不可
const Pi = 3.14

// 複数定数を同時に定義する
const (
	URL = "https://ueno"
	SiteName = "Portfolio"
)

// 同一値を定義する際の省略記法
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)
// 出力:1 1 1 A A A

// 変数はオーバーフローがあるが，定数はない


// 連続する整数を一気に定義する方法(0スタート)
const (
	i1 = iota
	i2
	i3
	i4
	i5
)
// 0 1 2 3 4

func main(){
	fmt.Println(A,B,C,D,E,F)

	fmt.Println(i1,i2,i3,i4,i5)


}
