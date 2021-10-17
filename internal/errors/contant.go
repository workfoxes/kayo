package errors

var (
	IndicatorNotFound = NotFoundError("Given indicator not found in kayo registory")
	BrokerNotFound    = NotFoundError("Given broker not found in kayo registory")
	RedisUnreachable  = RequestTimeoutError("Unable to reach redis, Please check the connection")
	MongoUnreachable  = RequestTimeoutError("Unable to reach MongoDB, please check the connection")
)
