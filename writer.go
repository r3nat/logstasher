package logstasher

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

// Writer wraps another writer and writes json events to it
type Writer struct {
	w      io.Writer
	fields map[string]interface{}
}

// NewWriter creates new wrapper around existing writer
func NewWriter(w io.Writer, fields map[string]interface{}) Writer {
	return Writer{
		w:      w,
		fields: fields,
	}
}

func (w Writer) Write(p []byte) (n int, err error) {
	e := map[string]interface{}{
		"@version":   1,
		"@timestamp": time.Now(),
		"message":    strings.TrimSpace(string(p)),
	}

	for k, v := range w.fields {
		e[k] = v
	}

	return len(p), json.NewEncoder(w.w).Encode(e)
}
