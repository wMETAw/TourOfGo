package main

import (
	"fmt"
	"math"
)

func isPrime(i int) bool {

	// 2,3は素数
	if i == 2 || i == 3 {
		return true
	}

	// 偶数なら素数ではない
	if i%2 == 0 {
		return false
	}

	// 元の数の平方数以下で割り切れる場合は素数ではない
	for n := 3; n <= int(math.Sqrt(float64(i))); n += 2 {
		if i%n == 0 {
			return false
		}
	}

	return true
}

func main() {

	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
