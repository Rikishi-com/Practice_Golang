# Golang

[GitHubリポジトリ](https://github.com/Rikishi-com/Practice_Golang)

---

[TOC]

## 実行方法

### 簡易実行方法
実行するソースコードに移動して以下のコマンドを打つ

```bash
go run main.go
```
- `main.go`は任意のファイル名

### 本番実行方法
```bash
go build main.go
```

```bash
./main   # macOS / Linux の場合
main.exe # Windows の場合
```

#### プロジェクト全体をビルドしたい場合
```bash
go build
```

#### 実行＋ビルドを一度にしたいとき
```bash
go run .
```
- `main.go`を含むディレクトリで実行
- これはあくまで一時的に実行するためのもの


---

### 実際の使用例

#### 1.基礎
```go
package main	//パッケージ宣言（パッケージ宣言は１つのみ）これのファイルの内容はmainパッケージになる

import (	//フォーマットパッケージ（標準パッケージ）
	"fmt"
	"time"
)	

func main(){	//エントリーポイント
	fmt.Println("Hello World")	//文字列出力
	fmt.Println(time.Now())
}
```

#### 2.変数(variant)

```go
package main

import "fmt"

// i5 := 500	//暗黙的な定義は関数外では定義不可
var i5 int = 500	//明示的な定義は関数外でも定義可能

func outer(){
	var s4 string = "ueno"
	fmt.Println(s4)
}

func main(){
	var i int = 100	//var 変数名 型 = 値
	fmt.Println(i)

	var s string = "Hello World"
	fmt.Println(s)

	var t, f bool = true, false	//複数変数を一度に設定する（同じ型の場合）
	fmt.Println(t,f)

	var(		//複数変数を一度に設定する（異なる型の場合）
		i2 int = 200
		s2 string = "Hello Go"
	)
	fmt.Println(i2,s2)

	var i3 int	//int型は初期値として0と設定される
	var s3 string	//string型は初期値として空文字が設定される
	fmt.Println(i3,s3)	//0と表示される

	i3 = 300
	s3 = "Hello Lang"
	fmt.Println(i3,s3)

	i = 150	//値の更新
	fmt.Println(i)

	//暗黙的な定義（型宣言が必要ない，しかし型がないわけではない．型推論で勝手に形を決めてくれる）
	i4 := 400	//変数名 := 値
	fmt.Println(i4)

	i4 = 450	//値の更新は暗黙的な定義でも同様の方法
	fmt.Println(i4)

	// i4 = "Hello"  //型推論によりi4はintとされているから，エラーが出る
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()	//outerは別関数

	// fmt.Println(s4)	//変数は関数内でのみ有効なのでエラーが出る

	// var s5 string = "yuki"	//Golangでは定義された変数は必ず使用する必要がある


}
```

#### 3.整数型(int)

```go
package main

import "fmt"

func main(){
	var i int = 100	//intの桁数は環境依存（使用環境が64bitで動いているならint64で動作する）

	var i2 int64 = 100	//明示的にすることも可能（int8,16,32,64などがある）

	fmt.Println(i + 50)

	// fmt.Println(i + i2)	
	//同じ整数型でもビット数が異なると演算はできない（基本的なPCは64bitだが，環境依存の64bitと明示的な64bitでは全く異なる型と判定される）

	fmt.Printf("%T\n", i)	//%Tは書式指定子．Tは型を表す

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))	//型変換は　型名(変数名)

}
```

#### ４.浮動小数点数型(float)

```go
package main

import "fmt"

func main(){
	var fl64 float64 = 1.1
	fmt.Println(fl64)

	fl := 2.2	//浮動小数点型の暗黙的な定義の場合，自動でfloat64になる（環境依存ではない）

	fmt.Println(fl)

	fmt.Printf("%T\n%T\n", fl,fl64)

	fmt.Println(fl + fl64)	//同じ型なので演算が可能（しかし丸め誤差が出る．これはどの言語でも一緒なので桁数制限すればOK）

	// var ueno float = 1.5	//floatという型はない

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)	//+inf = 正の無限大（infinity）

	ninf := - 1.0 / zero
	fmt.Println(ninf)	//-inf = 負の無限大(infinity)

	nan := zero / zero
	fmt.Println(nan)	//NaN = Not a Number = 定義できない

}
```

#### ５.論理データ型(boolean)

```go
package main

import "fmt"

func main(){
	var t, f bool = true, false
	fmt.Println(t,f)
}

```

#### ６.文字列型(string)

```go
package main

import "fmt"

func main(){
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	var s1 string = "300"	//もちろんstring
	fmt.Printf("%T\n", s1)

	//複数行はバッククォーテーションで囲む
	fmt.Println(`test	
	test
		test`)

	fmt.Println("\"")	//"を表示したいときは\を前にいれる

	fmt.Println("\\")	//バックスラッシュを表示するとき
}
```

#### ７.バイト型(byte)

```go
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
```


#### 8.配列型(array)

```go
package main

import "fmt"

func main(){
	var arr1 [3]int	//var 変数名　[要素数]要素の型名
	fmt.Println(arr1)	//[0 0 0]（整数型の初期値は0だから）
	fmt.Printf("%T\n", arr1)	//[3]intが型名となるため，Golangの配列は後から要素数を変更することができない

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)	//最後の要素は空文字

	arr3 := [3]int{1,2,3}	//暗黙的な宣言
	fmt.Println(arr3)

	arr4 := [...]string{"C","D"}	//[...]は与えられた要素数に自動で合わせてくれる
	fmt.Println(arr4)
	fmt.Printf("%T\n",arr4)	//型名は[2]stringになる（[...]stringではない）

	fmt.Println(arr2[0])	//値の取り出し（Pythonとかと変わらない）

	fmt.Println(arr3[3-1])	//要素数-1することで最後の値を取り出すことができる小技

	arr2[2] = "C"	//要素の更新

	fmt.Println(arr2)

	// var arr5 [4]int	
	// arr5 = arr1	//arr1とarr5の要素数が違うからエラーが出る
	// fmt.Println(arr5)

	fmt.Println(len(arr1))	//len関数は要素数を出力する（Pythonとかと一緒）
}
```

#### 9.インターフェース型(interface)

```go
package main

import "fmt"

func main(){
	var x interface{}	//すべての型と互換性を持つ特殊な型
	fmt.Println(x)	//nilが返る．nilはnoneと同じ意味

	x = 1
	fmt.Println(x)

	x = "ueno"
	fmt.Println(x)

	x = [3]int{1,2,3}	
	fmt.Println(x)

	// x = 2	//型特有の演算ができない
	// fmt.Println(x + 2)
}

```

#### 10.型変換

```go
package main

import (
	"fmt"
	"strconv"	//str convertの略
)

func main(){
	//変数名 = 型名(変換する変数名)

	var s string = "100"
	fmt.Println(s)
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi(s)	//文字列型から数値型への変換.　_の意味は値を破棄という意味．この関数は2つの値を返すが，1つしか使用しないため，破棄する
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	var i2 int = 100
	fmt.Println(i2)
	fmt.Printf("i2 = %T\n", i2)

	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)
}
```

#### 11.定数

```go
package main

import "fmt"

// 定数
// 頭文字を大文字にすることで，他のパッケージからも呼び出せる
// このパッケージ内だけなら，小文字でも良い
// 再定義不可
const Pi = 3.14

// 複数定数を同時に定義する
const (
	URL = "https://ueno"
	SiteName = "Portfolio"
)

// 同一値を定義する際の省略記法
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)
// 出力:1 1 1 A A A

// 変数はオーバーフローがあるが，定数はない


// 連続する整数を一気に定義する方法(0スタート)
const (
	i1 = iota
	i2
	i3
	i4
	i5
)
// 0 1 2 3 4

func main(){
	fmt.Println(A,B,C,D,E,F)

	fmt.Println(i1,i2,i3,i4,i5)


}
```

#### 12.演算子

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	// 加算演算子
	fmt.Println(1 + 2)

	// 減算演算子
	fmt.Println(5 - 1)

	// 乗算演算子
	fmt.Println(5 * 3)

	// 除算演算子
	fmt.Println(60 / 3)

	// 余り演算子
	fmt.Println(9 % 4)

	// 文字列連結演算子
	fmt.Println("ueno" + "yuki")

	n := 0
	// 加算代入演算子
	n += 4 //n = n + 4

	fmt.Println(n)

	// インクリメント演算子
	n++ //n = n + 1
	fmt.Println(n)

	// デクリメント演算子
	n-- //n = n - 1
	fmt.Println(n)

	s := "ABC"
	// 文字列連結代入演算子
	s += "DEF" //s = ABC + DEF

	fmt.Println(s)

	// べき乗（math.Pow を使用）
	// math パッケージをインポートする必要がある
	// math.Pow は float64 を返すので、整数計算に使う場合は注意
	// べき乗
	fmt.Println(math.Pow(2, 3))  // 2の3乗 = 8
	fmt.Println(math.Pow(9, 0.5)) // 平方根 = 3
}

```

#### 13.比較演算子

```go
package main

import "fmt"

func main() {

	// 等値演算子（等しい）
	fmt.Println(1 == 1) //true

	// 等値演算子（等しくない場合）
	fmt.Println(1 == 2) //false

	// 不等値演算子（等しくない）
	fmt.Println(1 != 2) //true

	// 不等値演算子（等しい場合）
	fmt.Println(1 != 1) //false

	// 小なり演算子
	fmt.Println(3 < 5) //true

	// 大なり演算子
	fmt.Println(5 > 3) //true

	// 小なりイコール演算子
	fmt.Println(4 <= 5) //true

	// 小なりイコール演算子（等しい場合）
	fmt.Println(5 <= 5) //true

	// 大なりイコール演算子
	fmt.Println(5 >= 4) //true

	// 大なりイコール演算子（等しい場合）
	fmt.Println(5 >= 5) //true

	// 文字列の比較
	fmt.Println("abc" == "abc") //true
	fmt.Println("abc" != "def") //true
}

```

#### 14.論理演算子

```go
package main

import "fmt"

func main() {
	// 論理AND演算子（&&）- 両方がtrueの場合のみtrue
	fmt.Println(true && true)   //true
	fmt.Println(true && false)  //false
	fmt.Println(false && true)  //false
	fmt.Println(false && false) //false

	// 論理OR演算子（||）- どちらかがtrueの場合true
	fmt.Println(true || true)   //true
	fmt.Println(true || false)  //true
	fmt.Println(false || true)  //true
	fmt.Println(false || false) //false

	// 論理NOT演算子（!）- 真偽値を反転
	fmt.Println(!true)  //false
	fmt.Println(!false) //true

	// 実際の条件での例
	age := 20
	hasLicense := true

	// AND演算子の実用例
	fmt.Println(age >= 18 && hasLicense) //true（18歳以上かつ免許あり）

	// OR演算子の実用例
	isWeekend := false
	isHoliday := true
	fmt.Println(isWeekend || isHoliday) //true（週末または祝日）

	// NOT演算子の実用例
	isWorking := false
	fmt.Println(!isWorking) //true（働いていない）

	// 複合条件の例
	fmt.Println((age >= 18 && hasLicense) || isHoliday) //複数の論理演算子を組み合わせ
}
```

#### 15.関数

```go
package main

import "fmt"

// 関数宣言
func Plus(x int, y int) int {
	return x + y
}

// 引数の型が同じなら省略可能
func Plus2(x , y int) int {
	return x + y
}

// 返り値が複数ある時
func Div (x,y int) (int,int) {
	q := x / y
	r := x % y
	return q, r
}

//func 関数名(引数 型...) 返り値の型 {return }

// 省略記法
func Double(price int) (result int){
	result = price * 2
	return
}

// 返り値のない関数
func Noreturn(){
	fmt.Println("No return")
	return
}

func main(){
	i := Plus(1,2)
	fmt.Println(i)
	// j, k := Div(2,4)
	// fmt.Println(j,k)


	// 片方しか値を使わない場合
	j, _ := Div(2,4)
	fmt.Println(j)

	x := Double(1000)
	fmt.Println(x)

	Noreturn()

}
```

#### 16.無名関数

```go
package main

import (
	"fmt"
	"math"
)

// 無名関数
// 初期化の即時実行・コールバック・goroutine内の
// 一時処理・外部変数を保持するクロージャなど、
// 局所的な処理をその場で手短にまとめたいときに使う

func main(){
	f := func(x, y int) int {
		return x + y
	}

	i := f(1,2)
	fmt.Println(i)

	//　そのまま値も入れてしまう省略記法
	i2 := func(p,q float64) float64 {
		return math.Pow(p,q)
	}(2,4)

	fmt.Println(i2)
}
```

#### 17.返り値が関数の関数

```go
package main

import "fmt"

// 関数を返す関数
// その場に合わせた柔軟な関数が作成できる

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

// xと引数を足し合わせる関数を作成する関数
// func(int) int: これは整数を受け取って整数を返す型という意味
func addr(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main(){
	f := ReturnFunc()
	f()

	g := addr(10)	// gは10と引数を足し合わせる関数
	fmt.Println(g(5))

	double_g := addr(20)	//double_gは20と引数を足し合わせる関数
	fmt.Println(double_g(10))
}
```

#### 18.引数が関数の関数

```go
package main

import "fmt"

// 関数を引数に取る関数

func CallFunction(f func()){
	f()
}

func main(){
	CallFunction(func(){
		fmt.Println("I'm a Ueno")
	})

}
```

#### 19.クロージャー

```go
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
```

#### 20.ジェネレーター

```go
package main

import "fmt"

// ジェネレーター
// 何らかの規則に従って連続した値を返す関数
// クロージャーの応用

func intergers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// xから始まるインクリメント
func argment_intergers(x int) func() int {
	i := x - 1
	return func() int {
		i++
		return i
	}
}

func main(){
	ints := intergers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	fmt.Println("---")

	arg_ints := argment_intergers(3)
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())

	fmt.Println("---")

	// 関数を新しく宣言したら新しくカウントする
	second_arg_ints := argment_intergers(3)
	fmt.Println(second_arg_ints())
	fmt.Println(second_arg_ints())
	fmt.Println(second_arg_ints())

	fmt.Println("---")

	// ここで使っても引き継がれる
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
	fmt.Println(arg_ints())
}
```
