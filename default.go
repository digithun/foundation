/*
 * Copyright (c) 2019. Octofox.io
 */

package foundation

import (
	"fmt"
	"os"
	"time"
)

func String(v string) *string {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func Int(v int) *int {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Time(t time.Time) *time.Time {
	return &t
}

func EnvStringOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Env %s not provide", key))
	}
	return value
}

func EnvString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
