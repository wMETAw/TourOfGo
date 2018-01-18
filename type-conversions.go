// 型変換
// 変数 v 、型 T があった場合、 T(v) は、変数 v を T 型へ変換します。
// C言語とは異なり、Goでの型変換は明示的な変換が必要です。
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
