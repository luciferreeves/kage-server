package logger

import (
	"io"
	"sync"
)

type LogLevel string

const (
	Debug   LogLevel = "DEBUG"
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARN"
	Error   LogLevel = "ERROR"
	Success LogLevel = "SUCCESS"
)

const (
	ColorReset   = "\033[0m"
	ColorCyan    = "\033[36m"
	ColorGray    = "\033[90m"
	ColorDebug   = "\033[90m"
	ColorInfo    = "\033[97m"
	ColorWarning = "\033[33m"
	ColorError   = "\033[31m"
	ColorSuccess = "\033[32m"
)

type Logger struct {
	prefix       string
	timestamp    bool
	timeFormat   string
	enableColors bool
	stdOutWriter io.Writer
	stdErrWriter io.Writer
	mu           sync.Mutex
}

type Option func(*Logger)
