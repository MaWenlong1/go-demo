package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("hello")
	var s student
	s.name = "12312"
	s.age = 12
	s.number = "11111"
	// s.print()
	var p Printer = s
	p.Prin()

	go task(1)
	go task(2)
	time.Sleep(time.Second * 3)

	done := make(chan bool)
	data := make(chan int)
	go consumer(data, done)
	go producer(data)

	<-done

	x, y := 1, 2
	x, y = y, x
	println(x, y)
}
func producer(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}
func consumer(data chan int, done chan bool) {
	for x := range data {
		println("receive:", x)
	}
	done <- true
}

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d:%d \n", id, i)
		time.Sleep(time.Second)
	}
}

type student struct {
	name   string
	age    int
	number string
}

func (s student) Prin() {
	fmt.Println(s)
}

type Printer interface {
	Prin()
}
