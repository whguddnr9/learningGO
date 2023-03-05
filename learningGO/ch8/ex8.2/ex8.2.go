package main

import "fmt"

func main () {
	const PI1 float64 = 3.1415926
	var PI2 float64 = 3.1415926

	//PI1 = 4 //여기서 오류 발생 (상수는 값 변경 x)
	PI2 = 4

	fmt.Printf("원주율: %f\n",PI1)
	fmt.Printf("원주율: %f\n",PI2)
	
}