package logus_test

import (
	"Samurai/internal/pkg/logus"
	murlog "github.com/Melenium2/Murlog"
	"testing"
	"time"
)

func MurlogConfig() *murlog.Config {
	m := murlog.NewConfig()
	m.CallerPref()
	m.TimePref(time.RFC1123)

	return m
}

func TestLogus_Log_LogMessage(t *testing.T) {
	l := logus.New(murlog.NewLogger(MurlogConfig()))
	l.Log("key", "value")
}

func TestLogus_LogMany_LogMoreMessages(t *testing.T) {
	l := logus.New(murlog.NewLogger(MurlogConfig()))
	l.LogMany(logus.NewLUnit("k", "v"), logus.NewLUnit("k1", "v1"))
}
