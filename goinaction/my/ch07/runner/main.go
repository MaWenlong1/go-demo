package main

import (
	"log"
	"os"
	"time"

	"github.com/goinaction/code/chapter7/patterns/runner"
)

const timeout = 3 * time.Millisecond

func main() {
	log.Println("开始工作")
	// 为本次执行分配超时时间
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("超时")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("中断")
			os.Exit(2)
		}
	}
	log.Println("结束工作")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
