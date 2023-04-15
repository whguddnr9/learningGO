package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 := stringheader.Data

	str += " World"
	addr2 := stringheader.Data

	str += " Welcome!"
	addr3 := stringheader.Data

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)

}

/*
go언어에서는 기존 문자열 메모리 공가늘 건드리지 않고 새로운 메모리 공간을 만들어서 두 문자열을 합치기 때문에 string합 연산 이후  주솟값이 변경된다.
따라서 문자열 불변 원칙이 준수가 된다.
*/
