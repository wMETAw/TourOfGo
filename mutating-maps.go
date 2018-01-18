// Mutating Maps
// map m の操作を見ていきましょう。
// m へ要素(elem)の挿入や更新:
//
// m[key] = elem
// 要素の取得:
//
// elem = m[key]
// 要素の削除:
//
// delete(m, key)
// キーに対する要素が存在するかどうかは、2つの目の値で確認します:
//
// elem, ok = m[key]
// もし、 m に key があれば、変数 ok は true となり、存在しなければ、 ok は false となります。
// なお、mapに key が存在しない場合、 elem はmapの要素の型のゼロ値となります。
//
// Note: もし elem や ok を宣言していないのであれば、次のように := で短く表現できます:
// elem, ok := m[key]
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
