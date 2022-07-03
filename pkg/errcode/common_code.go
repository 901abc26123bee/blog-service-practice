package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "Server Internal Error")
	InvalidParams             = NewError(10000001, "Invalid Params")
	NotFound                  = NewError(10000002, "NotFound")
	UnauthorizedAuthNotExist  = NewError(10000003, "Unauthorized, cannot find corresponding AppKey and AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "Unauthorized, token error")
	UnauthorizedTokenTimeout  = NewError(10000005, "Unauthorized, token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Unauthorized, token generation error")
	TooManyRequests           = NewError(10000007, "Too many requests")
)