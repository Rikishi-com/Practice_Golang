package main	//パッケージ宣言（パッケージ宣言は１つのみ）これのファイルの内容はmainパッケージになる

import (	//フォーマットパッケージ（標準パッケージ）
	"fmt"
	"time"
)	

func main(){	//エントリーポイント
	fmt.Println("Hello World")	//文字列出力
	fmt.Println(time.Now())
}