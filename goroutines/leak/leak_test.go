package leak_test

import (
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	go func() {
		for {
			<-time.After(time.Hour)
		}
	}()

}
