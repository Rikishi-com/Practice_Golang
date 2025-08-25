package main

import "fmt"

// マップ
// Pythonでいう辞書型(key_value)

func main(){
	// 明示的宣言
	var m = map[string]int{"A":100,"B":200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"A":100, "B":200}
	fmt.Println(m2)

	m3 := map[int]string{
		0:"JAPAN",
		1:"USA",
		100:"Thai",
	}

	fmt.Println(m3[100])	// Thai

	m4 := make(map[string]int)	// 初期値のみ

	fmt.Println(m4)

	s,e := m3[0]
	s2,e2 := m3[2]	// 第二引数:boolean(引数があればTrue,なければFalse)

	fmt.Println(s,e)
	fmt.Println(s2,e2)

	m3[1] = "US"
	fmt.Println(m3[1])	// US

	// delete関数
	
	delete(m3,1)
	fmt.Println(m3)	//delete(マップ名，インデックス要素)

	

}