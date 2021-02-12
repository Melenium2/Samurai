package logus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog_SendLog(t *testing.T) {
	logurl := "http://146.0.36.96:7491/log"

	log := &Log{
		Type:       "error",
		Module:     "samurai",
		Message:    "Hello",
	}
	assert.NoError(t, log.SendLog(logurl))
}
