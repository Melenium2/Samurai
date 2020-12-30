package retry

import (
	"math"
	"math/rand"
	"time"
)

// struct represent implementation of exponential and linear backoff
type backoff struct {
	maxAttempts  uint
	attemptNum   uint
	factor       float64
	intensity    time.Duration
	maxRetryTime time.Duration
	timeNow      time.Duration
	isLinear     bool
}

// IsEnd check if count of attempts more then max
func (b *backoff) IsEnd() bool {
	return b.attemptNum > b.maxAttempts
}

// Next attempt
func (b *backoff) Next() {
	b.attemptNum++
}

// Calculate linear of exponential NextDelay.
//Base on factor, intensity and attempts
func (b *backoff) NextDelay() time.Duration {
	if b.isLinear {
		b.timeNow = time.Duration(float64(b.intensity) * float64(b.attemptNum))
		if b.timeNow > b.maxRetryTime {
			b.timeNow = b.maxRetryTime
		}
		return b.timeNow
	}

	rand.Seed(time.Now().UnixNano())
	attemptNow := math.Max(float64(b.attemptNum), 1)
	bf, max := float64(b.intensity), float64(b.maxRetryTime)
	for bf < max && attemptNow > 0 {
		bf *= b.factor * attemptNow
		attemptNow--
	}

	bf *= 1 + 0.2*(rand.Float64()*2-1)

	if bf > max {
		bf = max
	}

	b.timeNow = time.Duration(bf)
	return b.timeNow
}

// Return AttemptNow
func (b *backoff) AttemptNow() int {
	return int(b.attemptNum)
}

// Return TimeNow
func (b *backoff) TimeNow() time.Duration {
	return b.timeNow
}

// Reset backoff
func (b *backoff) Reset() {
	b.attemptNum = 0
	b.timeNow = 0
}

// Creates NewBackoff
func NewBackoff(maxAttempts uint, factor float64, maxRetryTime, intensity time.Duration, isLinear ...bool) *backoff {
	linear := false
	if len(isLinear) > 0 {
		linear = isLinear[0]
	}
	return &backoff{
		maxAttempts:  maxAttempts,
		factor:       factor,
		intensity:    intensity,
		maxRetryTime: maxRetryTime,
		isLinear:     linear,
		attemptNum:   1,
	}
}
