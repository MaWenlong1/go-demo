package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	// clicktime 默认连点时间 intervaltime 默认连点间隔
	var clicktime, intervaltime int
	flag.IntVar(&clicktime, "c", 20, "连点持续时间（单位s），默认为20s")
	flag.IntVar(&intervaltime, "i", 1, "连点间隔，默认为1")
	flag.Parse()
	clicktime = clicktime * 1000
	fmt.Println("=============程序开始！=============")
	// 是否处于点击状态
	isClick := false
	// 是否结束点击
	isEnd := false
	registerFunc("q", func(e hook.Event) {
		fmt.Println("\n退出！")
		hook.End()
	})
	registerFunc("w", func(e hook.Event) {
		go func() {
			if isClick {
				return
			}
			fmt.Println("\n连点开始！")
			isClick = true
			isEnd = false
			for i := 0; i < clicktime/intervaltime && !isEnd; i++ {
				robotgo.MouseClick("left", true)
				time.Sleep(time.Duration(intervaltime) * time.Millisecond)
			}
			isClick = false
			isEnd = true
		}()
	})
	registerFunc("e", func(e hook.Event) {
		if !isClick {
			return
		}
		fmt.Println("\n连点结束！")
		isClick = false
		isEnd = true
	})
	s := hook.Start()
	<-hook.Process(s)
}

func registerFunc(key string, f func(hook.Event)) {
	hook.Register(hook.KeyDown, []string{key}, f)
}
