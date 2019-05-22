package foundation

import "time"

func String(v string) *string {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func Int(v int) *int {
	return &v
}

func Time(t time.Time) *time.Time {
	return &t
}
