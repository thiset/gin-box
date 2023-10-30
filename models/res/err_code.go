package res

type ErrorCode int

const (
	SystemError ErrorCode = 1001 // 系统错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SystemError: "系统错误",
	}
)
