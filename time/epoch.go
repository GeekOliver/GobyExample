package main

import (
	"fmt"
	"time"
)

func main()  {
	now := time.Now()

	secs := now.Unix()

	nanos := now.UnixNano()

	fmt.Println(now)

	// 注意 `UnixMillis` 是不存在的，所以要得到毫秒数的话，
	// 你要自己手动的从纳秒转化一下。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 你也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
