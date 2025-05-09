package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func NewLogger(options ...Option) *Logger {
	l := &Logger{
		timeFormat:   time.RFC3339,
		enableColors: true,
		timestamp:    false,
		stdOutWriter: os.Stdout,
		stdErrWriter: os.Stderr,
	}

	for _, opt := range options {
		opt(l)
	}

	return l
}

func (l *Logger) WithPrefix(prefix string) *Logger {
	l.prefix = prefix
	return l
}

func (l *Logger) WithSubPrefix(subPrefix string) *Logger {
	var prefix string
	if l.prefix != "" {
		prefix = fmt.Sprintf("%s:%s", l.prefix, subPrefix)
	} else {
		prefix = subPrefix
	}

	l.prefix = prefix
	return l
}

func (l *Logger) WithTimestamp() *Logger {
	l.timestamp = true
	return l
}

func (l *Logger) WithoutTimestamp() *Logger {
	l.timestamp = false
	return l
}

func (l *Logger) WithTimeFormat(format string) *Logger {
	l.timeFormat = format
	return l
}

func (l *Logger) WithColors() *Logger {
	l.enableColors = true
	return l
}

func (l *Logger) WithoutColors() *Logger {
	l.enableColors = false
	return l
}

func (l *Logger) getLevelColor(level LogLevel) string {
	stringBuilder := func(level LogLevel) string {
		return fmt.Sprintf("%s%s%s", l.getMessageColor(level), l.getLevelString(level), ColorReset)
	}

	switch level {
	case Debug:
		return stringBuilder(Debug)
	case Info:
		return stringBuilder(Info)
	case Warning:
		return stringBuilder(Warning)
	case Error:
		return stringBuilder(Error)
	case Success:
		return stringBuilder(Success)
	default:
		return stringBuilder(level)
	}
}

func (l *Logger) getLevelString(level LogLevel) string {
	switch level {
	case Debug:
		return l.padString("DEBUG", 7)
	case Info:
		return l.padString("INFO", 7)
	case Warning:
		return l.padString("WARN", 7)
	case Error:
		return l.padString("ERROR", 7)
	case Success:
		return l.padString("SUCCESS", 7)
	default:
		return string(level)
	}
}

func (l *Logger) getMessageColor(level LogLevel) string {
	switch level {
	case Debug:
		return ColorDebug
	case Info:
		return ColorInfo
	case Warning:
		return ColorWarning
	case Error:
		return ColorError
	case Success:
		return ColorSuccess
	default:
		return ColorReset
	}
}

func (l *Logger) padString(str string, length int) string {
	if len(str) >= length {
		return str
	}
	padding := length - len(str)
	return str + strings.Repeat(" ", padding)
}

func (l *Logger) log(level LogLevel, message any, fatal bool) {
	var builder strings.Builder

	if l.timestamp {
		if l.enableColors {
			builder.WriteString(ColorGray)
		}
		builder.WriteString(time.Now().Format(l.timeFormat))
		if l.enableColors {
			builder.WriteString(ColorReset)
		}
		builder.WriteString(" ")
	}

	if l.enableColors {
		builder.WriteString(l.getLevelColor(level))
	} else {
		builder.WriteString(l.getLevelString(level))
	}
	builder.WriteString(" ")

	if l.prefix != "" {
		if l.enableColors {
			builder.WriteString(ColorCyan)
		}
		builder.WriteString(l.padString(fmt.Sprintf("[%s]", l.prefix), 15))
		if l.enableColors {
			builder.WriteString(ColorReset)
		}
		builder.WriteString(" ")
	}

	if l.enableColors {
		builder.WriteString(l.getMessageColor(level))
	}

	switch msg := message.(type) {
	case error:
		builder.WriteString(msg.Error())
	case string:
		builder.WriteString(msg)
	default:
		builder.WriteString(fmt.Sprintf("%v", msg))
	}

	if l.enableColors {
		builder.WriteString(ColorReset)
	}
	builder.WriteString("\n")

	output := builder.String()

	l.mu.Lock()
	defer l.mu.Unlock()

	if level == Error || level == Warning {
		l.stdErrWriter.Write([]byte(output))
	} else {
		l.stdOutWriter.Write([]byte(output))
	}

	if fatal {
		os.Exit(1)
	}
}

func (l *Logger) Debug(message any) {
	l.log(Debug, message, false)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.log(Debug, fmt.Sprintf(format, args...), false)
}

func (l *Logger) Info(message any) {
	l.log(Info, message, false)
}

func (l *Logger) Infof(format string, args ...any) {
	l.log(Info, fmt.Sprintf(format, args...), false)
}

func (l *Logger) Warn(message any) {
	l.log(Warning, message, false)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.log(Warning, fmt.Sprintf(format, args...), false)
}

func (l *Logger) Error(message any) {
	l.log(Error, message, false)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.log(Error, fmt.Sprintf(format, args...), false)
}

func (l *Logger) Fatal(message any) {
	l.log(Error, message, true)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.log(Error, fmt.Sprintf(format, args...), true)
}

func (l *Logger) Success(message any) {
	l.log(Success, message, false)
}

func (l *Logger) Successf(format string, args ...any) {
	l.log(Success, fmt.Sprintf(format, args...), false)
}
