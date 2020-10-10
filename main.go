package main

import (
	"GobyExample/Basic"
)

func main() {
	Basic.Sum(10, 20, 30, 50)

	nums := []int{1,2,3,4}
	Basic.Sum(nums...)
}