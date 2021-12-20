package backoff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackoff(t *testing.T) {
	b := Bc{
		Config: DefaultConfig,
	}
	d := b.Backoff(10)
	assert.NotEmpty(t, d)
}
