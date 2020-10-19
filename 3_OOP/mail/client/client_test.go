package client

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("john.doe@localhost", "Sample", "Hello, World!")
}