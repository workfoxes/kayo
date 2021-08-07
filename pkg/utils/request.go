package utils

import (
	"net/http"
	"time"
)

// Const vars for rate limiter
const (
	DefaultMaxRequestJobs   int32 = 50
	DefaultMaxRetryAttempts       = 3
	DefaultMutexLockTimeout       = 50 * time.Millisecond
	drainBodyLimit                = 100000
	proxyTLSTimeout               = 15 * time.Second
	userAgent                     = "User-Agent"
)

// Vars for rate limiter
var (
	MaxRequestJobs   = DefaultMaxRequestJobs
	MaxRetryAttempts = DefaultMaxRetryAttempts
)

// Requester : is the struct for the request client
type Requester struct {
	HTTPClient *http.Client
	//limiter            Limiter
	Name       string
	UserAgent  string
	maxRetries int
	jobs       int32
	//Nonce              nonce.Nonce
	disableRateLimiter int32
	backoff            Backoff
	//retryPolicy        RetryPolicy
	//timedLock *timedmutex.TimedMutex
}
