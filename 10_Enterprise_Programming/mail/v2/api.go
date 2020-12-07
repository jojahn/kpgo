package mail

type Sender interface {
	Send(Mail)
}

type Mail struct {
	Subject string
	Topic string
	Content string
}

func NewSender() *LogSender {
	return &LogSender{}
}