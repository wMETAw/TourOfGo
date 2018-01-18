// struct (構造体)は、フィールド( field )の集まりです。
//( type 宣言は型を宣言するためのものです)
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
