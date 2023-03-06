package main

import "fmt"

func main() {
	var age = 22

	if age >= 10 && age <= 15 { //age가 10이상 15이하
		fmt.Println("You're age young")
	} else if age > 30 || age < 20 { //age 가 20대가 아닐때
		fmt.Println("you are not 20s")
	} else { //age 가 20 대 일때
		fmt.Println("Best age of your life")
	}
}
