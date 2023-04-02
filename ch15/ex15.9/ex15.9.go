package main

import "fmt"

func main() {
	str := "Hello 월드!"

	for _, v := range str {
		fmt.Printf("Type: %T Value: %d charater: %c\n", v, v, v)
	}
}
