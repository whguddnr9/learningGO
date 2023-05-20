package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	/* 아래와 같이 선언하면 for 구문으로 반복해야하는 단점이 생김
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}
	*/

	slice2 := append([]int{}, slice1...)

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
