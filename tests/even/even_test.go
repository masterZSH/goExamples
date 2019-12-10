package even

import (
	"testing"
)

func TestEven(t *testing.T) {
	if Even(7) {
		t.Log("7 not even")
		t.Error()
	}
	if !Even(4) {
		t.Log("4 is even")
		t.Error()
	}
}

func TestOdd(t *testing.T) {
	if !Odd(7) {
		t.Log("7 is odd")
		t.Error()
	}
	if Odd(4) {
		t.Log("4 is not odd")
		t.Error()
	}
}
