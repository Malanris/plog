package log

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func _zeroTime() time.Time {
	return time.Time{}
}

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithTimeFunction(_zeroTime),
		WithNoStyles())
	cases := []struct {
		name     string
		expected string
		msg      string
		kvs      []interface{}
		f        func(msg interface{}, kvs ...interface{})
	}{
		{
			name:     "simple message",
			expected: "INFO info\n",
			msg:      "info",
			kvs:      nil,
			f:        logger.Info,
		},
		{
			name:     "ignored message",
			expected: "",
			msg:      "this is a debug message",
			kvs:      nil,
			f:        logger.Debug,
		},
		{
			name:     "message with keyvals",
			expected: "INFO info key1=val1 key2=val2\n",
			msg:      "info",
			kvs:      []interface{}{"key1", "val1", "key2", "val2"},
			f:        logger.Info,
		},
		{
			name:     "error message with keyvals",
			expected: "ERROR info key1=val1 key2=val2\n",
			msg:      "info",
			kvs:      []interface{}{"key1", "val1", "key2", "val2"},
			f:        logger.Error,
		},
	}
	for _, c := range cases {
		buf.Reset()
		t.Run(c.name, func(t *testing.T) {
			c.f(c.msg, c.kvs...)
			assert.Equal(t, c.expected, buf.String())
		})
	}
}

func TestOffLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf), WithTimeFunction(_zeroTime),
		WithNoStyles(), WithLevel(OffLevel))
	cases := []struct {
		name     string
		expected string
		msg      string
		kvs      []interface{}
		f        func(msg interface{}, kvs ...interface{})
	}{
		{
			name:     "simple message",
			expected: "",
			msg:      "error",
			kvs:      nil,
			f:        logger.Error,
		},
	}
	for _, c := range cases {
		buf.Reset()
		t.Run(c.name, func(t *testing.T) {
			c.f(c.msg, c.kvs...)
			assert.Equal(t, c.expected, buf.String())
		})
	}
}

func TestLogHelper(t *testing.T) {
	var buf bytes.Buffer
	logger := New(WithOutput(&buf),
		WithCaller(), WithNoStyles())

	helper := func() {
		logger.Helper()
		logger.Info("helper func")
	}

	helper()
	_, file, line, ok := runtime.Caller(0)
	require.True(t, ok)
	assert.Equal(t, fmt.Sprintf("INFO <log/%s:%d> helper func\n", filepath.Base(file), line-1), buf.String())
}
