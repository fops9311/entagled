package entangle

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
