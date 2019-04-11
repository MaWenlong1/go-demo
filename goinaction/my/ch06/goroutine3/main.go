package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	court := make(chan int)
	wg.Add(2)
	go player("张三", court)
	go player("李四", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s 胜利\n", name)
			return
		}
		n := rand.Intn(200)
		if n%13 == 0 {
			fmt.Printf("%s 失败\n", name)
			close(court)
			return
		}
		fmt.Printf("%s 击球数：%d \n", name, ball)
		ball++

		court <- ball
	}
}
