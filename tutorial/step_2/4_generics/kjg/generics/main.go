package main

import "fmt"

// Number 인터페이스: int64와 float64를 포함
type Number interface {
	int64 | float64
}

// SumNumbers 함수: Number 제약 조건을 사용하는 제네릭 함수
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
			s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
			"first":  34,
			"second": 12,
	}
	floats := map[string]float64{
			"first":  35.98,
			"second": 26.99,
	}

	fmt.Printf("타입 제약을 사용한 제네릭 합계: %v, %v\n",
			SumNumbers(ints),
			SumNumbers(floats))
}
