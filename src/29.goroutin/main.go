package main

import (
	"fmt"
	"time"
)

// ゴルーチン
// 並行処理

// func sub(){		// これだとmain　loopにたどり着かない（上から処理するから，無限ループのため抜け出せない）
// 	for {
// 		fmt.Println("sub loop")
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main(){
// 	sub()
// 	for {
// 		fmt.Println("main loop")
// 		time.Sleep(200 * time.Microsecond)
// 	}

// }

func sub(){
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go sub()	// goいれるだけで並行処理可能
	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Millisecond)
	}

}
