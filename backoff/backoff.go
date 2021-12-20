package backoff

import (
	"math/rand"
	"time"
)

type Config struct {
	// BaseDelay is the amount of time to backoff after the first failure.
	BaseDelay time.Duration
	// Multiplier is the factor with which to multiply backoffs after a
	// failed retry. Should ideally be greater than 1.
	Multiplier float64
	// Jitter is the factor with which backoffs are randomized.
	Jitter float64
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

var DefaultConfig = Config{
	BaseDelay:  1.0 * time.Second,
	Multiplier: 1.6,
	Jitter:     0.2,
	MaxDelay:   120 * time.Second,
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type Bc struct {
	Config Config
}

// Backoff 失败重试时间
// retries 失败次数
func (bc Bc) Backoff(retries int) time.Duration {
	if retries == 0 {
		return bc.Config.BaseDelay
	}
	b := float64(bc.Config.BaseDelay)
	max := float64(bc.Config.MaxDelay)

	// 计算到最大 每次乘以Multiplier
	for b < max && retries > 0 {
		b *= bc.Config.Multiplier
		retries--
	}

	// 超过最大值 重置为最大值
	if b > max {
		b = max
	}

	// 随机
	// r.Float64() [0, 1)
	b *= 1 + bc.Config.Jitter*(r.Float64()*2-1)

	return time.Duration(b)
}
