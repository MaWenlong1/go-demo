package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)
	// wg 用来等待程序完成
	// 技术加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("开始goroutine")

	// 声明一个匿名函数并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'z'+1; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()
	// 声明一个匿名函数并创建一个goroutine
	go func() {
		// 在函数退出时调用Done来通知main函数工作完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'Z'+1; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()

	// 等待goroutine结束
	fmt.Println("等待线程完成")
	wg.Wait()
	fmt.Println("程序结束")
}
