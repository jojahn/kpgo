package timeout

import (
	"fmt"
	"time"
)

func setTimeout(callback func() int, timeout time.Duration) (int, error) {
	channel := time.After(timeout)
	callbackChannel := make(chan int)
	go func() {
		callbackChannel <- callback()
	}()
	select {
	case <-channel:
		return 0, fmt.Errorf("Timeout has run out!")
	case res := <-callbackChannel:
		return res, nil
	}
}