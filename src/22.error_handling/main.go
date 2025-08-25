package main

import (
	"fmt"
	"strconv"
)

func main(){
	// var s string = "100"
	var s string = "A"

	i, err := strconv.Atoi(s)	// この関数にはエラーの場合，2つ目の変数にエラーの内容が入る．
	if err != nil {		// エラーではない場合，err = nil になるので，この条件
		fmt.Println(err)
	}
	fmt.Printf("%T\n",i)
}
