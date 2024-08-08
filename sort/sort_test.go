package sort

import (
	"testing"
)

func TestFib(t *testing.T) {

	t.Run("test3 ", func(t *testing.T) {
		if got:=Fib(100);got!=2 {
			t.Failed()
		}
	})
	t.Run("test3 recursion", func(t *testing.T) {
		if got:=FibRecur(100);got!=2 {
			t.Failed()
		}
	})
}
