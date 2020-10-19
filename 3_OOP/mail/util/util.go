package util

import (
	"github.com/jojahn/kpgo/3_OOP/mail/smtp"
)

type Registry struct {
	services map[string]interface{}
}

func NewRegistry() Registry {
	services := make(map[string]interface{})
	return Registry{services}
}

func (r *Registry) Get(locator string) interface{} {
	if r.services[locator] != nil {
		return r.services[locator]
	}
	if locator == "mail.Sender" {
		var sender interface{} = smtp.SenderImpl{}
		r.services[locator] = sender
		return sender
	}
	return nil
}