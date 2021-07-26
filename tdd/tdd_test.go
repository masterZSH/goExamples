package tdd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	Greet(os.Stdout, "test")

	var writer bytes.Buffer
	actual := "test"
	Greet(&writer, actual)
	assert.Equal(t, writer.String(), actual)
}
