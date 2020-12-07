package client

import "github.com/jojahn/kpgo/10_Enterprise_Programming/mail"

func test() {
	message := mail.Message{}
	sender := mail.LogSender{}
	sender.Send()
}
