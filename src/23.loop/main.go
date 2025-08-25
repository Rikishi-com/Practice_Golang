package main

import "fmt"

// 繰り返し処理
// Golangにはwhile文が存在しないのでfor文を使う

func main(){
	// for {	// 無限ループ
	// 	fmt.Println("Infinity_Loop")
	// }
	
	i :=0
	for {
		i++;
		if i == 3{
			break	// for文を強制終了するもの（一番近いfor文を終了させる）
		}
		fmt.Println("2_Loop")
	}

	point := 0
	for point < 10 {
		fmt.Println("10_Loop")
		point++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---")

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue	// この条件のときにここの時点で頭に戻る
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break	// この条件のときに終了する
		}
		fmt.Println(i)
	}

	fmt.Println("---")

	// 配列を使ったループ
	arr := [3]int{4,5,6}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 範囲式
	for i, v := range arr {
		fmt.Printf("インデックス番号:%d\n配列要素:%d\n", i,v )	// いらない要素は_使う
	}

	fmt.Println("---")

	// スライスを使ったもの
	sl := []string{"Python", "Go", "TypeScript"}

	for i, s := range sl {
		fmt.Printf("インデックス番号:%d\n配列要素:%s\n", i, s )
	}

	fmt.Println("---")

	// map
	m := map[string]int{"apple":100, "banana":200}

	for k, v := range m {
		fmt.Println(k,v)
	}


}