package main

import "fmt"

type notifier interface {
	notify()
}
type notifier1 interface {
	notify()
	notify1()
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("user----->name:%s\t email:%s\n", u.name, u.email)
}
func (u *user) notify1() {
	fmt.Printf("--------user----->name:%s\t email:%s\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (u *admin) notify() {
	fmt.Printf("admin----->name:%s\t email:%s\n", u.name, u.email)
}

func main() {
	lisa := user{"lisa", "lisa@gmail.com"}
	super := admin{"super", "super@gmail.com"}
	send(&lisa)
	send(&super)
}
func send(n notifier) {
	n.notify()
}
