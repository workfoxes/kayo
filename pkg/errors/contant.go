package errors

var (
	RedisUnreachable = RequestTimeoutError("Unable to reach redis, Please check the connection")
	MongoUnreachable = RequestTimeoutError("Unable to reach MongoDB, please check the connection")
)
