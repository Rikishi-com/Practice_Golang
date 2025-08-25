package main

import "fmt"

// map
// for
func main(){
	m := map[int]string{
		0: "ueno",
		1: "yuki",
	}

	for k,v := range m {
		fmt.Println(k,v)
	}
}