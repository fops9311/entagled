package main

import "fmt"

func main() {
	in, out := entagle()
	in(1)
	a := out().(int)
	a++
	fmt.Println(a)
	fmt.Println(out())
	in(false)
	fmt.Println(out())
}
func entagle() (in func(interface{}), out func() interface{}) {
	var data interface{}
	in = func(d interface{}) {
		data = d
	}
	out = func() interface{} {
		return data
	}
	return in, out
}
