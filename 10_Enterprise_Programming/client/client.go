package client

import "github.com/jojahn/kpgo/10_Enterprise_Programming/mail/v2"

type Client struct {}

func (c *Client) notifyAdmin(alert string) {
	message := mail.Mail{
		"admin@localhost",
		"Notify",
		alert,
	}
	sender := mail.LogSender{}
	sender.Send(message)
}
