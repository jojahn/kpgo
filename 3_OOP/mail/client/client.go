package client

import (
	"github.com/jojahn/kpgo/3_OOP/mail/mail"
	"github.com/jojahn/kpgo/3_OOP/mail/util"
	)

var Registry = util.NewRegistry()

func SendMail(to string, subject string, text string) error {
	// Create an implementation for the smtp.Sender interface
	var sender = Registry.Get("mail.Sender").(mail.Sender)
	email := mail.Message{To: to, Subject: subject, Text: text}
	return sender.Send(email)
}