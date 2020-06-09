package os

import (
	"testing"
)

func TestGetArgs(t *testing.T) {
	args := getArgs()
	t.Fatal(args)
	if len(args) != 0 {
	}
}


