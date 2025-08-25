package main

import "fmt"


// ラベル付きfor
// break,continueなどの後にラベル名をつけるとそこまでの範囲のfor文すべてに影響を与える
func main(){
	i := 0
	OutSideLoop:
	for {
		fmt.Println("2回")
		i++
		if (i == 2){
			break OutSideLoop
		}
		InSideLoop:
		for{
			for{
				fmt.Println("一回のみ")
				break InSideLoop
			}
		}
	}
}
