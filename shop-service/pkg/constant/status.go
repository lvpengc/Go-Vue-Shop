// Package constant status.go 定义状态码
package constant

const (
	// Succ 成功状态码
	Succ int = 0
)

// web api 的错误状态码
const (
	PostErr     int = iota + 1001 // post错误，一般显示单一的错误msg
	PostErrForm                   // 表单错误，一般显示一组错误msg
)

// Auth Token 的错误状态码
const (
	InvalidToken int = iota + 2001 // 无效token
	ExpiredToken                   // 过期token
)

// Casbin 权限错误状态码
const (
	NOPERMISSION int = iota + 3001 // casbin校验无权限
)
