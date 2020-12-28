package retry_test

import (
	"Samurai/internal/pkg/retry"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRetry_ShouldReturnErrorWith3Ticks(t *testing.T) {
	err := retry.Go(func() error {
		return fmt.Errorf("error))))")
	}, retry.WithMaxAttempts(3), retry.WithMaxRetryTime(time.Second * 10), retry.WithFactor(1.6))

	assert.Error(t, err)
}

func TestRetry_ShouldReturnResultRightAfterFirstTry(t *testing.T) {
	var res string
	err := retry.Go(func() error {
		res = "123"
		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, "123", res)
}

func TestRetry_ShouldReturnResultOnLastTry(t *testing.T) {
	var res string
	counter := 0
	err := retry.Go(func() error {
		if counter != 2 {
			counter++
			return fmt.Errorf("not yet")
		}
		res = "123"

		return nil
	}, retry.WithMaxAttempts(3))

	assert.NoError(t, err)
	assert.Equal(t, "123", res)
	assert.Equal(t, 2, counter)
}
