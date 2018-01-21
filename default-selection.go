// time.Tick(<Duration>) は 指定した間隔で時刻を配信する同期チャネル を返す関数
// tick := time.Tick(time.Millisecond)した場合のtickは1ms間隔で時刻を送信する
// 通常のチャネル同様に<-tickで受ける
// チャネルなのでselect case文やfor t := range tick { ... }文で受信処理が可能

// ２つのTime型を比べて未来かどうかを判定します。
//
// t, err := time.Parse(timeformat, "2013-04-09 22:57:14")
// f, err := time.Parse(timeformat, "2012-01-01 22:57:14")
// fmt.Println(f.After(t))
// After()を使います。あるいはBefore()
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}

	}
}
