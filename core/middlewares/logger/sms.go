package logger

import (
	"core/config"
)

type Sms struct {
	Conf *config.Sms
}

func NewSms(conf *config.Sms) (*Sms, error) {
	s := &Sms{}
	s.Conf = conf
	return s, nil
}
