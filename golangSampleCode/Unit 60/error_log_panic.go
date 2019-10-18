package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	if n == 1 {                           // 1일때만
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러를 리턴
}

func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)   // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Panic(err) // 에러 문자열이 출력되고 런타임 에러 발생
	}

	// 런타임 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
