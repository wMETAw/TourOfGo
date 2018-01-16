package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand.Seed(time.Now().UnixNano())
// という箇所で乱数の元になるシードを与えるのがポイント。
// なぜ、シードを与えるかというと、ここで使っている関数rand.Intnは
// randパッケージに新しい乱数ジェネレータ生成するときに
// 同じ初期化 rand.NewSource(1) をしているそうで、
// シードをあたえて、異なる乱数を出す必用があるからだそうです。
func main() {

	// 1 7 7 9 1 8 5 0 6 0 4 1 2 9 8 4 1 5 7 6になる
	for i := 0; i < 20; i++ {
		fmt.Print(rand.Intn(10), " ")
	}
	fmt.Println(" ")

	// randをNewSourceで初期化して実行
	my_rand := rand.New(rand.NewSource(1))
	for i := 0; i < 20; i++ {
		fmt.Print(my_rand.Intn(10), " ")
	}
	fmt.Println(" ")

	// シードを与えて実行
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		fmt.Print(rand.Intn(10), " ")
	}
}
