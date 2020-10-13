package Limit

import (
	"fmt"
	"net/http"
	"time"
	"golang.org/x/time/rate"
)


func TestRateLimit()  {
	r := rate.Every(1 *time.Millisecond)//每一毫秒投放一次令牌
	limit := rate.NewLimiter(r, 10)//桶的容量为10
	
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        if limit.Allow() {
            fmt.Printf("请求成功，当前时间：%s\n", time.Now().Format("2006-01-02 15:04:05"))
        } else {
            fmt.Printf("限流了\t")
        }
    })

	_ = http.ListenAndServe(":8088", nil)
}

