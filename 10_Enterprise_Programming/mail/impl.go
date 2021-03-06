package mail

import (
	log "github.com/sirupsen/logrus"
)

type LogSender struct {
	Sender
}

func (s *LogSender) Send(message Message) {
	log.WithFields(log.Fields{
		"subject": message.Subject,
		"topic": message.Topic,
	}).Info(message.Content)
}
