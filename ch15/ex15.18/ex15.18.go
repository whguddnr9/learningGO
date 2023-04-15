package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a')) // 합 연산 사용
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder // string Builder 변수 선언
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a')) // string.Builder 사용
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}

/* string builder을 사용하면 위와같이 문자열을 변경시킬 수 있게된다. string Builder을 사용하지 않고 문자열을 계속 복사해서 변경하게 된다면 메모리 낭비로 이어질 수 있다.*/
