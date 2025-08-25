package main

import "fmt"

func main(){
	var H,W,N int
	fmt.Scan(&H,&W,&N)

	var stp []string
	var tmp string

	for i := 0;i < (H * N); i++ {
		fmt.Scan(&tmp)
		stp = append(stp, tmp)
	}

	var R,C,temp int
	fmt.Scan(&R,&C)

	for k := 0; k < R * C ; k++{
		fmt.Scan(&temp)
		tmp := (temp - 1)
		fmt.Print(stp[temp - 1])
		if ((k + 1) % C) == 0 {
			fmt.Print("\n")
		}

	}
}
