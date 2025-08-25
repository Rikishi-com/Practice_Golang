package main

import "fmt"

func init(){	// 最初に実行される関数（初期化）
	fmt.Println("init")
}

func init(){	// 何個でも作れる．その場合は上から実行（しかしあまり作る意味はない）
	fmt.Println("init2")
}

func main(){
	fmt.Println("main")
}
