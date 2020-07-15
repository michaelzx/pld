package pld_errs

import "net/http"

// Forbidden 用户未授权
type Forbidden struct {
	HttpErr
}

func NewForbidden(code int, message string) *Forbidden {
	r := new(Forbidden)
	r.Status = http.StatusForbidden
	r.Code = http.StatusForbidden
	r.Message = message
	return r
}

func (e *Forbidden) Error() string {
	return e.Message
}
