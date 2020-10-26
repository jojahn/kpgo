package util

import (
	"github.com/jojahn/kpgo/3_OOP/mail/smtp"
)

type Registry struct {
	services map[string]interface{}
}

func NewRegistry() Registry {
	services := make(map[string]interface{})
	services["mail.Sender"] = smtp.SenderImpl{}
	return Registry{services}
}

func (r *Registry) Get(locator string) interface{} {
	return r.services[locator]
}