package main

import "fmt"

func main() {
	dan := 2

	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 20 {
				break
			}
		}
		b = 1
		dan++
		if dan == 20 {
			break
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
