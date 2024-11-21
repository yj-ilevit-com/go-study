package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Reverse 함수: 문자열을 UTF-8 단위로 반전
func Reverse(s string) (string, error) {
    // UTF-8 유효성 검사
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }

    // UTF-8 문자열을 rune 단위로 처리
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
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
