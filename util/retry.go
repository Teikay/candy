package util

import (
	"time"
)

const (
	// defaultRetryAttemptMax Max retry attempts.
	defaultRetryAttemptMax = 20
	// defaultRetryDurationMin is the initial duration, 100 ms
	defaultRetryDurationMin = 100
	// defaultRetryDurationMax is the max amount of duration, 2000 ms
	defaultRetryDurationMax = 2000
)

//Retry 自带sleep的重试.
type Retry struct {
	attempts    int
	attemptMax  int
	durationMin int64
	durationMax int64
}

// RetryOption retry option.
type RetryOption func(*Retry)

//RetryAttemptMax 设置最大重试次数, 0 为一直重试
func RetryAttemptMax(v int) RetryOption {
	return func(r *Retry) { r.attemptMax = v }
}

//RetryDurationMin 最小间隔时间
func RetryDurationMin(v int) RetryOption {
	return func(r *Retry) { r.durationMin = int64(v) }
}

//RetryDurationMax 最小间隔时间
func RetryDurationMax(v int) RetryOption {
	return func(r *Retry) { r.durationMax = int64(v) }
}

// NewRetry new retry.
func NewRetry(ops ...RetryOption) *Retry {
	r := &Retry{
		attemptMax:  defaultRetryAttemptMax,
		durationMin: defaultRetryDurationMin,
		durationMax: defaultRetryDurationMax,
	}
	for _, op := range ops {
		op(r)
	}
	return r
}

//Valid check max attempts
func (r *Retry) Valid() bool {
	if r.attemptMax != 0 && r.attempts >= r.attemptMax {
		return false
	}
	return true
}

//Attempts current attempt times.
func (r *Retry) Attempts() int {
	return r.attempts
}

// Next wait sleep.
func (r *Retry) Next() {
	dur := time.Now().UnixNano()%(r.durationMin<<uint32(r.attempts)) + r.durationMin
	if dur > r.durationMax {
		dur = r.durationMax
	}
	sleep := time.Duration(dur) * time.Millisecond
	time.Sleep(sleep)
	r.attempts++
}
