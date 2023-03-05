package main

import "fmt"

func Divide(a,b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a/b, true
}

func main () {
	c,success := Divide(9,3)	// := 를 이용해서 각각 대입 방법
	fmt.Println(c,success)
	d, success := Divide(9,0)
	fmt.Println(d,success)
}