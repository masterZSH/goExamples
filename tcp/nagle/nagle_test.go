package nagle

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	go Start()
	time.Sleep(2 * time.Second)

	conn, err := net.Dial("tcp", "localhost:6666")
	assert.Nil(t, err)
	content := make([]byte, 1024*4)
	for k := range content {
		content[k] = 1
	}

	for i := 0; i < 10; i++ {
		conn.Write(content)
	}
}
