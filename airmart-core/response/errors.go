package response

import (
	"fmt"
)

func Err(code int, msg string) error {
	return &Codes{
		Code: code,
		Msg:  msg,
	}
}

func (c *Codes) SetMsg(msg string) *Codes {
	c.Msg = msg
	return c
}

func (c *Codes) Msgf(template string, args ...interface{}) *Codes {
	c.Msg = fmt.Sprintf(template, args...)
	return c
}

// Err todo 可能改成c.LogErr = err
func (c *Codes) Err(err error) *Codes {
	c.Msg = err.Error()
	return c
}

func (c *Codes) Error() string {
	return fmt.Sprintf("Code: %v, Msg: %v", c.Code, c.Msg)
}
