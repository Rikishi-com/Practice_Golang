package main

import "fmt"

func main(){
	ch1 := make(chan int,5)

	ch1 <- 1

	close(ch1)	// クローズすると受信しなくなるだけ，値は消えない

	// fmt.Println(ch1)	//エラーが起こる

	i, ok := <-ch1	// クローズしても値は出せる

	fmt.Println(i,ok)
}