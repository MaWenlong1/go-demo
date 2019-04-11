package main

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"Bill", "bill@gmail.com"}
	send(&u)
}

func send(n notifier) {
	n.notify()
}
