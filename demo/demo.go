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
	slice := []int{1, 2, 3, 4, 5}
	new := slice[1:3]
	new = append(new, 22222)
	fmt.Println(new)
	fmt.Println(slice)
	new = append(new, 33333)
	fmt.Println(new)
	fmt.Println(slice)
	new = append(new, 44444)
	fmt.Println(new)
	fmt.Println(slice)
	new = append(new, 55555)
	fmt.Println(new)
	fmt.Println(slice)
}
