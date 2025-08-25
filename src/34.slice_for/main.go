package main

import "fmt"

func main(){
	sl := []string{"A","B","C"}

	for i, v := range sl {	// i:インデックス番号 , v:配列要素
		fmt.Println(i,v)
	}

	// 古典的
	for i := 0; i < len(sl); i++{
		fmt.Println(i,sl[i])
	}
}