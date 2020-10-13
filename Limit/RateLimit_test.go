package Limit

import (
	"fmt"
	"net/http"
	"testing"
)



func getApi()  {
	api := "http://localhost:8088/"
	res ,err := http.Get(api)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Printf("get api success\n")
	}
}
//go test -bench=. -benchtime=3s -run=none
func Benchmark_Main(b *testing.B)  {
	for i:= 0; i < b.N; i++ {
		getApi()
	}
}