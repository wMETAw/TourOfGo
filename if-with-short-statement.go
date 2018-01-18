// if ステートメントは、 for のように、条件の前に、評価するための簡単なステートメントを書くことができます。
// ここで宣言された変数は、 if のスコープ内だけで有効です。
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
