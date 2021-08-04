package main

import "fmt"

func main() {
	in, out := entagle()
	in(true)
	fmt.Println(out())
	in(false)
	fmt.Println(out())
}
func entagle() (in func(bool), out func() bool) {
	var data bool
	in = func(d bool) {
		data = d
	}
	out = func() bool {
		return data
	}
	return in, out
}
