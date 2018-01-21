// チャネルが、goroutine間で通信するための素晴らしい方法であることを見てきました。
//
// しかし、通信が必要ない場合はどうでしょう？ コンフリクトを避けるために、
// 一度に1つのgoroutineだけが変数にアクセスできるようにしたい場合はどうでしょうか？
//
// このコンセプトは、排他制御( mutual exclusion )と呼ばれ、このデータ構造を指す一般的な名前は mutex (ミューテックス)です。
//
// Goの標準ライブラリは、排他制御をsync.Mutexと次の二つのメソッドで提供します:
//
// Lock
// Unlock
// Inc メソッドにあるように、 Lock と Unlock で囲むことで排他制御で実行するコードを定義できます。
//
// Value メソッドのように、mutexがUnlockされることを保証するために defer を使うこともできます。
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
