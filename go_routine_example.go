// 루틴: 쓰레드 비슷?
// 채널: 데이터 공유?
// 쓰레드 비슷?

package main

import (
	"fmt"
	"time"
)

func work(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "working", i, "hours")
		time.Sleep(time.Second)
	}
}

func main() {
	work("programmer")
	work("designer")
	work("producer")
	work("marketer")

	go work("programmer")
	go work("designer")
	go work("producer")
	go work("marketer")

	wait := 0
	fmt.Scanln(&wait)

}
