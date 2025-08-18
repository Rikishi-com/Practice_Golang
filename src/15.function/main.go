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
