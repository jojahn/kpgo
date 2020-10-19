package mail

import "fmt"

type Message struct {
	To string
	Subject string
	Text string
}

type Sender interface {
	// Send a smtp to a given address with a subject and text.
	Send(message Message) error
}

func (m Message) String() string {
	return fmt.Sprintf("Message { To %s, Subject %s, Text %s }",
		m.To, m.Subject, m.Text)
}