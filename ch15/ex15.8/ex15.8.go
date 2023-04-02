package main

import "fmt"

func main() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Type: %T Value: %d charater: %c\n", arr[i], arr[i], arr[i])
	}
}
