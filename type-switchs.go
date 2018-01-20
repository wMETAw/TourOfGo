// 型switch はいくつかの型アサーションを直列に使用できる構造です。
//
// 型switchは通常のswitch文と似ていますが、
// 型switchのcaseは型(値ではない)を指定し、
// それらの値は指定されたインターフェースの値が保持する値の型と比較されます。
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I donot know about type %T\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
