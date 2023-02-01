package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b) // 입력 두 개 받기
	if err != nil {            // 에러가 발생하면 에러코드 출력
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b) // 정상 입력되면 입려값 출력
	}
}
