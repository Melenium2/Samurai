package retry_test

import (
	"Samurai/internal/pkg/retry"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var standartBackoff = retry.NewBackoff(10, 2, time.Second*10, time.Millisecond*500)

func TestBackoff_Next(t *testing.T) {
	standartBackoff.Next()

	assert.Equal(t, 1, standartBackoff.AttemptNow())

	for i := 0; i < 5; i++ {
		standartBackoff.Next()
	}

	assert.Equal(t, 6, standartBackoff.AttemptNow())
}

func TestBackoff_IsEnd(t *testing.T) {
	for i := 0; i < 11; i++ {
		standartBackoff.Next()
	}

	assert.True(t, standartBackoff.IsEnd())
}

func TestBackoff_NextDelay(t *testing.T) {
	tt := []struct {
		name       string
		linear     bool
		maxTime    time.Duration
		intensity  time.Duration
		iterations int
		shouldbe   time.Duration
	}{
		{
			name:       "test linear time function. should be 700 on TimeNow check",
			linear:     true,
			maxTime:    time.Second,
			intensity:  time.Millisecond * 100,
			iterations: 6,
			shouldbe:   time.Millisecond * 700,
		},
		{
			name: "test linear function. should be max time value",
			linear: true,
			maxTime: time.Second,
			intensity: time.Millisecond * 100,
			iterations: 20,
			shouldbe: time.Second,
		},
		{
			name:       "test exponential time function. may be random time less then 5 seconds",
			linear:     false,
			maxTime:    time.Second * 10,
			intensity:  time.Millisecond * 400,
			iterations: 2,
			shouldbe:   time.Second * 5,
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			boff := retry.NewBackoff(10, 2, test.maxTime, test.intensity, test.linear)

			for i := 0; i < test.iterations; i++ {
				boff.Next()
			}

			d := boff.NextDelay()
			assert.GreaterOrEqual(t, d.Seconds(), test.shouldbe.Seconds())
		})
	}
}
