package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 给每个可用的核心分配一个逻辑处理器
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)

	go printPrime("A")
	go printPrime("B")

	wg.Wait()

	fmt.Println("程序退出")

}
func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 50000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("完成", prefix)
}
