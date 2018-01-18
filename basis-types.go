// Go言語の基本型(組み込み型)は次のとおりです:
//
// uint8       符号なし  8-ビット 整数 (0 to 255)
// uint16      符号なし 16-ビット 整数 (0 to 65535)
// uint32      符号なし 32-ビット 整数 (0 to 4294967295)
// uint64      符号なし 64-ビット 整数 (0 to 18446744073709551615)
//
// int8        符号あり  8-ビット 整数 (-128 to 127)
// int16       符号あり 16-ビット 整数 (-32768 to 32767)
// int32       符号あり 32-ビット 整数 (-2147483648 to 2147483647)
// int64       符号あり 64-ビット 整数 (-9223372036854775808 to 9223372036854775807)
//
// float32     IEEE-754 32-ビット 浮動小数値
// float64     IEEE-754 64-ビット 浮動小数値
//
// complex64   float32の実数部と虚数部を持つ複素数
// complex128  float64の実数部と虚数部を持つ複素数
//
// byte        uint8の別名(エイリアス)
// rune        int32の別名(エイリアス)
// (訳注：runeとは古代文字を表す言葉(runes)ですが、Goでは文字そのものを表すためにruneという言葉を使います。)
//
// 例では、いくつかの型の変数を示しています。また、変数宣言は、インポートステートメントと同様に、まとめて( factored )宣言可能です。
//
// int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bitです。
// サイズ、符号なし( unsigned )整数の型を使うための特別な理由がない限り、整数の変数が必要な場合は int を使うようにしましょう。
//
// golang complex(複素数)型を使う
// https://qiita.com/intelf___/items/039eccffd422321ec6dd
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	// %T	この値の型をGo言語の構文で表現する
	// %v	デフォルトフォーマットを適用した値 構造体を出力する際、+フラグ(%+v)を加えるとフィールド名が表示される
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
