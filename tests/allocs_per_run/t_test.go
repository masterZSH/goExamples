package allocs

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	n := testing.AllocsPerRun(100, func() {
		f1()
	})
	fmt.Println(n)
}
