package main

import "plugin"
import "github.com/jojahn/kpgo/10_Enterprise_Programming/mail"

type Sender interface {
	Send(interface{})
}

func main() {
	p, err := plugin.Open("../plugin/plugin.so")
	if err != nil {
		panic(err)
	}

	newSenderSymbol, err := p.Lookup("NewSender")
	if err != nil {
		panic(err)
	}
	var NewSender func() *mail.Sender = newSenderSymbol.(func() *mail.Sender)
	var sender mail.Sender = *NewSender()
	message := mail.Message{"abc", "def", "Hello, World!"}
	&sender.Send(message)
}