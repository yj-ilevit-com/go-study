package main

import (
	"fmt"
)

// Reverse 함수: 문자열을 반전
func Reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"

	// Reverse 호출 시 반환값 두 개를 처리
	rev, revErr := Reverse(input)
	if revErr != nil {
			fmt.Printf("Error reversing input: %v\n", revErr)
			return
	}

	doubleRev, doubleRevErr := Reverse(rev)
	if doubleRevErr != nil {
			fmt.Printf("Error reversing reversed input: %v\n", doubleRevErr)
			return
	}

	// 결과 출력
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
