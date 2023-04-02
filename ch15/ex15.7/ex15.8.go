package main

import "fmt"

func main() {
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("Type: %T Value: %d charater: %c\n", str[i], str[i], str[i])
	}
}
