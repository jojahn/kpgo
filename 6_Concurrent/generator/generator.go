package generator

func fib() <-chan int {
	fibChan := make(chan int)
	go func() {
		i, j := 0, 1
		for {
			i, j = j, i + j
			fibChan <- i
		}
	}()
	return fibChan
}