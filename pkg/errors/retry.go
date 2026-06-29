package errors

import "time"

type retry struct {
	delay time.Duration
}

func newRetryAfter(d time.Duration) retry {
	return retry{delay: d}
}
