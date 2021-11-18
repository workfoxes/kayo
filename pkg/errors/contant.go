package errors

import (
	_errors "github.com/workfoxes/calypso/pkg/errors"
)

var (
	RedisUnreachable = _errors.RequestTimeoutError("Unable to reach redis, Please check the connection")
	MongoUnreachable = _errors.RequestTimeoutError("Unable to reach MongoDB, please check the connection")
)
