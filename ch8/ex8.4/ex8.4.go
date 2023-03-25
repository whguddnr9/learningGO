package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100      // 상수 타입 자동변환됨
	var b int = FloatPI * 100 // 상수는 float 형식임으로 오류 발생

	fmt.Println(a)
	fmt.Println(b)
}

/*
# ex8.4
./ex8.4.go:10:14: cannot use FloatPI * 100
(constant 314 of type float64) as int value in variable declaration
*/
