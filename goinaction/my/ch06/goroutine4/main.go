package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()
}
func Runner(baton chan int) {
	var newRunner int
	runner := <-baton

	fmt.Printf("%d 开始跑步\n", runner)
	time.Sleep(100 * time.Millisecond)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("%d 到达接力线\n", newRunner)
		go Runner(baton)
	} else {
		fmt.Printf("%d 结束比赛\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("%d 将接力棒交给%d\n", runner, newRunner)
	baton <- newRunner
}
