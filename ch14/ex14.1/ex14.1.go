package main

import "fmt"

func main() {
	a := 500
	var p *int

	p = &a

	fmt.Printf("value of p : %p\n", p)
	fmt.Printf("memory value pointed by p : %d\n", *p)

	*p = 100
	fmt.Printf("value of a :%d\n", a)
}
