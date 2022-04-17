package dataconv

import "fmt"

func ShowConv()  {
	var a = 24
	var b = 2.0
	c := float64(a) * b
	fmt.Println(c)

	// string으로 변환할때 아주 좋음.
	precision := fmt.Sprintf("%.2f",b)
	// %T = typeof
	fmt.Printf("%s - %T\n", precision , precision)
}