package main

import "fmt"

type user struct {
	name       string
	email      string
	age        int
	privileged bool
}
type admin struct {
	person user
	level  string
}

func main() {
	var bill user
	fmt.Println(bill)
	lisa := user{
		name:       "lisa",
		email:      "3333@gmail.com",
		age:        20,
		privileged: true,
	}
	fmt.Println(lisa)
	fred := admin{
		person: user{"lisa", "3333@gmail.com", 20, true},
		level:  "super",
	}
	fmt.Println(fred)
}
