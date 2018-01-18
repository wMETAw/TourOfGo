// 明示的な型を指定せずに変数を宣言する場合( := や var = のいずれか)、変数の型は右側の変数から型推論されます。
// 右側の変数が型を持っている場合、左側の新しい変数は同じ型になります:
package main

import "fmt"

func main() {
	v := 42 // change me
	fmt.Printf("v is of type %T\n", v)
}
