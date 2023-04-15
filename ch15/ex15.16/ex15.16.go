package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceheader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data)
	fmt.Printf("slice:\t%x\n", sliceheader.Data)

}

/* string 의 특정 문자만 변경하고 싶을때에는 slice로 타입변환을 해야한다.
이는 새로운 변수를 할당한 것으로 각 변수는 똑같은 변수를 가르키고 있지 않다.*/
