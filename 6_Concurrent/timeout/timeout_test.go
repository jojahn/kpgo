package timeout

import (
	"testing"
	"time"
	"fmt"
)

func TestTimeout(t *testing.T) {
	res, err := setTimeout(func() int {
		time.Sleep(10 * time.Millisecond)
		return 1
	}, 1*time.Second)
	
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("operation returned %d", res)
	}
}

