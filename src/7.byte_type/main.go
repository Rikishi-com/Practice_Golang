package main

import "fmt"

func main(){
	byteA := []byte{72,73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))	//文字コードを実際の文字に変換

	c := []byte("HI")
	fmt.Println(c)	//文字コードに変換

	fmt.Println(string(c))
}
