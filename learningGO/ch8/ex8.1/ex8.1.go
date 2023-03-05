package main

import "fmt"

func main () {
	const C int = 10 //상수 선언

	var b int = C * 20

	C = 20 // Error 상수는 다른값으로 변환 x
	fmt.Println(&C)  //Error 상수는 주소 x 
}

/*
# ex8.1
./ex8.1.go:8:6: b declared and not used
./ex8.1.go:10:2: cannot assign to C (constant 10 of type int)
./ex8.1.go:11:15: invalid operation: cannot take address of C (constant 10 of type int)
*/ 