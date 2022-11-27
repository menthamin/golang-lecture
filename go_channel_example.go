// 루틴: 쓰레드 비슷?
// 채널: 데이터 공유?
// 쓰레드 비슷?

package main

import (
	"fmt"
	"time"
)

func work(s string, done chan string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, "working", i, "hours")
		time.Sleep(time.Millisecond * 500)
	}
	done <- s // 채널에 s를 보내는 것
}

func main() {
	done := make(chan string, 8)
	work("programmer", done)
	work("designer", done)
	work("producer", done)
	work("marketer", done)

	go work("programmer", done)
	go work("designer", done)
	go work("producer", done)
	go work("marketer", done)

	// wait := <-done
	for n := range done {
		fmt.Println((n))
	}
	// fmt.Println(wait)
	fmt.Scanln()
}
