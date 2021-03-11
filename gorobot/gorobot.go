package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	// fmt.Println("Press any key")
	// hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
	// 	fmt.Println(e)
	// })
	// s := hook.Start()
	// <-hook.Process(s)
	add()
	fmt.Println("======")
	// 	Nanosecond Duration = 1
	// 	Microsecond = 1000 * Nanosecond
	// 	Millisecond = 1000 * Microsecond
	// 	Second = 1000 * Millisecond
	// 	Minute = 60 * Second
	// 	Hour = 60 * Minute
	// time.Sleep(10 * time.Second)
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	robotgo.EventHook(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		robotgo.EventEnd()
	})

	fmt.Println("--- Please press w---")
	robotgo.EventHook(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		go func() {
			for i := 0; i < 1000; i++ {
				robotgo.MouseClick(`left`, true)
				time.Sleep(time.Duration(50) * time.Millisecond)
			}
		}()
	})

	fmt.Println("--- Please press q---")
	robotgo.EventHook(hook.MouseWheel, []string{}, func(e hook.Event) {
		fmt.Println("press q---")
		robotgo.EventEnd()
	})
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
func low() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		fmt.Println("hook: ", ev)
	}
}
func event() {
	ok := robotgo.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	keve := robotgo.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}

	mleft := robotgo.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}
func test() {
	/* ========================= 按键操作 ======================== */
	// 向上滚动：3行
	robotgo.ScrollMouse(3, `up`)
	// 向下滚动：2行
	robotgo.ScrollMouse(2, `down`)

	// 按下鼠标左键
	// 第1个参数：left(左键) / center(中键，即：滚轮) / right(右键)
	// 第2个参数：是否双击
	robotgo.MouseClick(`left`, false)

	// 按住鼠标左键
	robotgo.MouseToggle(`down`, `left`)
	// 解除按住鼠标左键
	robotgo.MouseToggle(`up`, `left`)

	/* ========================= 位置操作 ======================== */
	// 将鼠标移动到屏幕 x:800 y:400 的位置（闪现到指定位置）
	robotgo.MoveMouse(800, 400)

	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	robotgo.MoveMouseSmooth(800, 400)

	// 将鼠标移动到屏幕 x:800 y:400 的位置（模仿人类操作）
	// 第3个参数：纵坐标x 的延迟到达时间
	// 第4个参数：横坐标y 的延迟到达时间
	robotgo.MoveMouseSmooth(800, 400, 20.0, 200.0)

	/* ========================= 组合操作 ======================== */
	// 移动鼠标到 x:800 y:400 后，双击鼠标左键
	robotgo.MoveClick(800, 400, `left`, true)

	/* ========================= 所在位置 ======================== */
	// 获取当前鼠标所在的位置
	x, y := robotgo.GetMousePos()
	print(`x：`, x, ` y：`, y)
}
