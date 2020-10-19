package smtp

import (
	"fmt"
	"github.com/jojahn/kpgo/3_OOP/mail/mail"
)

type SenderImpl struct {}

func (s SenderImpl) Send(msg mail.Message) error {
	fmt.Println(fmt.Sprintf("Sending Message using SMTP: %s", msg))
	return nil
}
