package main

import "fmt"

// 定义一个用户类型
type user struct {
	name  string
	email string
}

// notify使用值接受者实现一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail使用指针接受者实现一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// user类型的值可以调用
	// 使用值接受者声明的方法
	// 使用指针接受声明的方法
	bill := user{"BILL", "bill@gmail.com"}
	bill.notify()
	bill.changeEmail("bill@change.com")
	bill.notify()
	// user类型的值的指针可以调用
	// 使用值接受者声明的方法
	// 使用指针接受声明的方法
	lisa := &user{"LISA", "lisa@gmail.com"}
	lisa.notify()
	lisa.changeEmail("lisa@change.com")
	lisa.notify()
}
