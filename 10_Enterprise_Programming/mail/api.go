package mail

type Sender interface {
	Send(Message)
}

type Message struct {
	Subject string
	Topic string
	Content string
}