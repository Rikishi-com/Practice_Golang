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
