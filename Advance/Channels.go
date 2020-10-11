package Advance

import "fmt"

func TestChannel()  {
	ms := make(chan string)
	defer close(ms)
	go func() {
		ms <- "ping"
	}()

	msg := <- ms
	fmt.Println(msg)
}
