package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t { // range를 사용하면 i에는 배열의 크기 v 에는 배열 값이 저장된다!!
		fmt.Println(i, v)
	}
}
