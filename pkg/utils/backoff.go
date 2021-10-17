package utils

import "time"

// Backoff determines how long to wait between request attempts.
type Backoff func(n int) time.Duration

// DefaultBackoff is a default strategy for backoff after a retryable request failure.
func DefaultBackoff() Backoff {
	return LinearBackoff(100*time.Millisecond, time.Second)
}

// LinearBackoff applies a backoff increasing by a base amount with each retry capped at a maximum duration.
func LinearBackoff(base, max time.Duration) Backoff {
	return func(n int) time.Duration {
		d := base * time.Duration(n)
		if d > max {
			return max
		}

		return d
	}
}
