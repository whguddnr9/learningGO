package main

import "fmt"

func chageArray(array2 [5]int) {
	array2[2] = 200
}

func chageSlice(slice2 []int) {
	slice2[2] = 200
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	chageArray(array)
	chageSlice(slice)

	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
}
