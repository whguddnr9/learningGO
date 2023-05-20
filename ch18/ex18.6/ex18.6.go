package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5) //slice1에 4,5 추가 -> 기존 배열크기보다 더  큰값

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 //slice1의 [1]에 100으로 변경

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500) //500 추가

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

}

/*
slice에서 기존 배열의 크기보다 더 큰 값이 저장되어야 할때는 새로운 배열을 만들게되고
위와같이 slice1과 slice2가 같은 배열을 바라보고 있어도 한 변수에서 더 큰값이 저장되면
서로다른 배열을 가르키게 된다.
*/
