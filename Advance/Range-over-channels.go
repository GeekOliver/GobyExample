package Advance

import "fmt"

func TestRangeOverChannels()  {
	queue := make(chan string, 2)
	queue <- "1"
	queue <- "2"

	close(queue)


	for e := range queue {
		fmt.Println(e)
	}
}
