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
