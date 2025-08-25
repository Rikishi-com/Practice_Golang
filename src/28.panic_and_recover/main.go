package main

import "fmt"

// panic recover
// 例外処理
// あまり推奨されない（ランタイムを強制終了するため）
// 例外処理では，エラーハンドリング(err)が推奨される

func main(){
	defer func() {
		if x := recover(); x != nil {	// panicから回復するためのもの，panicは強制終了のため，実質deferとの組み合わせでのみ動作する
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("No Written")
}