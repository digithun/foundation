/*
 * Copyright (c) 2019. Inception Asia
 * Maintain by DigithunWorldwide ❤
 * Maintainer
 * - rungsikorn.r@digithunworldwide.com
 * - nipon.chi@digithunworldwide.com
 * - mai@digithunworldwide.com
 */

package logger

type Fields map[string]interface{}

const defaultLoggerName = "solar"

var defaultLogger = New(defaultLoggerName)
var Default = defaultLogger

func WithField(key string, value interface{}) *Logger {
	return defaultLogger.WithField(key, value)
}

func WithError(err error) *Logger {
	return defaultLogger.WithError(err)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}
func Println(args ...interface{}) {
	defaultLogger.Println(args...)
}
func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}
