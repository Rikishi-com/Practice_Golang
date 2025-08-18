package main

import "fmt"

// クロージャの強み：
// 1) 「状態」を関数の外に漏らさずに保持できる（カプセル化）
// 2) 「ふるまい」を部分適用・カスタマイズして返せる（関数を返す）
// 3) 計算結果のキャッシュ（メモ化）で高速化できる

// 一個表示が遅れる関数（これはあまり強みは活かせてない）
// 無名関数で扱ったもののほうが強みを活かせている
func Later() func(string) string{
	var store string	// 最初は空文字
	return func(next string) string{
		s := store	// sにstoreを代入
		store = next	// storeにnextを代入
		return s	// s(中身はstore)を返す
	}
}

func main(){
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("World"))
}