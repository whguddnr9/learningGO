package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	// slice = append(slice, 0)

	idx := 2
	/*
		반복문으로
		for i := len(slice) - 2; i >= idx; i-- {
			slice[i+1] = slice[i]
		}

		slice[idx] = 100
	*/

	/*
		append 이중 중첩으로
		slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	*/

	// copy를 이용하여 이중중첩으로 인한 메모리 낭비 해결
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100

	fmt.Println(slice)
}
