package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v \n", v)
	fmt.Println(v.String())
	fmt.Println(v.Type())

	max := math.MaxInt64
	fmt.Println(reflect.TypeOf(max))
	fmt.Println(max)
}
