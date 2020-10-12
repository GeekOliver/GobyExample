package Limit

import (
	"fmt"
	"time"
)

func TestTimers()  {

	 // 定时器表示在未来某一时刻的独立事件。你告诉定时器
    // 需要等待的时间，然后它将提供一个用于通知的通道。
    // 这里的定时器将等待 2 秒。
	t1 := time.NewTimer(time.Second * 2)
	// `<-timer1.C` 直到这个定时器的通道 `C` 明确的发送了
    // 定时器失效的值之前，将一直阻塞。
	<-t1.C
	fmt.Println("t1 1 expired")

	 // 如果你需要的仅仅是单纯的等待，你需要使用 `time.Sleep`。
    // 定时器是有用原因之一就是你可以在定时器失效之前，取消这个
    // 定时器。这是一个例子
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}