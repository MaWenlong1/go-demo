package main

import (
	"fmt"
)

func getPrime(num int) {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for i := 2; i < num; i++ {
		origin <- num
	}
	close(origin)
	<-wait
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}

func main() {
	getPrime(1000)
}
