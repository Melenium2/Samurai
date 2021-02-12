package logus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Log struct {
	Type       string `json:"type,omitempty"`
	Module     string `json:"module,omitempty"`
	Message    string `json:"message,omitempty"`
	StackTrace string `json:"stack_trace,omitempty"`
}

func (l Log) String() string {
	return fmt.Sprintf("type=%s module=%s message=%s stacktrace=%s", l.Type, l.Module, l.Message, l.StackTrace)
}

func (l *Log) SendLog(logurl string) error {
	by, err := json.Marshal(l)
	if err != nil {
		return err
	}
	resp, err := http.Post(logurl, "application/json", bytes.NewReader(by))
	if err != nil {
		return err
	}
	if resp.StatusCode > 200 {
		return ErrLogRequest
	}

	return nil
}
