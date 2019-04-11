package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("停止")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

// 检测之前的shutdown标志来决定是否提前终止
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("%s 开始工作\n", name)
		time.Sleep(250 * time.Millisecond)

		// 是否要停止工作
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("%s 停止\n", name)
			break
		}
	}
}
