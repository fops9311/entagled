package entangled

import (
	"testing"
)

func TestEntangle(t *testing.T) {
	in, out := Entangle()
	in(1)
	a := out().(int)
	a++
	if a != 2 {
		t.Errorf("in(1) --> out().(int)++ = %d; want 2", a)
	}
	if out().(int) != 1 {
		t.Errorf("in(1) --> out().(int) = %d; want 1", a)
	}
	in(false)
	if out().(bool) != false {
		t.Errorf("in(false) --> out().(bool) = %d; want false", a)
	}
}
func TestToWay(t *testing.T) {
	in, out := ToWay()
	in.emiter(1)
	result := out.reciever().(int)
	if result != 1 {
		t.Errorf("in.emiter(1) --> out.reciever().(int) = %d; want 1", result)
	}
	out.emiter(1)
	result = in.reciever().(int)
	if result != 1 {
		t.Errorf("out.emiter(1) --> in.reciever().(int) = %d; want 1", result)
	}
}
