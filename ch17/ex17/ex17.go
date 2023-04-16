/*
	숫자 마추기 게임 만들기

	- 랜덤한 숫자값을 저장하기
		숫자를 완벽한 랜덤값을 만들기 위해
	- 숫자 입력받기
	- 숫자 비교하기


*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 난수 생성 함수. 랜드 씨드 값을 신간 값으로 초기화 해서 난수 생성.
func RandNum() int {
	rand.Seed(time.Now().UnixMicro())

	n := rand.Intn(100)
	// fmt.Println(n)

	return n
}

// 입력으로 들어온 값을 버퍼에 저장한다. 이 저장한 값으로 오류를 검사한다.
var stdin = bufio.NewReader(os.Stdin)

// 값을 입력받고 오류를 확인하기 위한 함수
func InputValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	//오류 발생시 버퍼를 초기화 한다. 초기화를 위해 오류 후 뒤따라오는 줄바꿈을 읽는다.
	if err != nil {
		stdin.ReadString('\n') // 버퍼를 이때 줄바꿈을 읽지 않을시 중복으로 오류를 인식한다.
	}
	return n, err
}

func main() {
	quiz := RandNum()
	count := 0
	for {
		fmt.Println("숫자값을 입력하세요.")
		num, err := InputValue()
		if err != nil {
			fmt.Println("숫자만 입력해 주세요.")
		} else {
			count++
			fmt.Printf("입력하신 숫자는 %d 입니다.", num)
			if num < quiz {
				fmt.Printf("입력한 숫자가 더 작습니다. 시도한 횟수 %d \n", count)
			} else if num > quiz {
				fmt.Printf("입력한 숫자가 더 큽니다. 시도한 횟수 %d \n", count)
			} else {
				fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수 %d \n", count)
				break
			}
		}
	}
}
