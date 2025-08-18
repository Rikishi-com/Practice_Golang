package main

import "fmt"

func main(){
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	var s1 string = "300"	//もちろんstring
	fmt.Printf("%T\n", s1)

	//複数行はバッククォーテーションで囲む
	fmt.Println(`test	
	test
		test`)

	fmt.Println("\"")	//"を表示したいときは\を前にいれる

	fmt.Println("\\")	//バックスラッシュを表示するとき
}
