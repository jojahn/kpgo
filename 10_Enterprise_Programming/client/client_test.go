package client

import "testing"

func TestClient(t *testing.T) {
	client := Client{}
	client.notifyAdmin("Hello, World!")
}