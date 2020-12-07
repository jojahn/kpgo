package mail

type Sender interface {
	Send(Message)
}

type Message struct {
	subject string
	topic string
	content string
}