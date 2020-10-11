package Basic

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func TestLoop()  {
	fmt.Println(fact(1))
	fmt.Println(fact(2))
	fmt.Println(fact(3))
	fmt.Println(fact(4))
	fmt.Println(fact(5))
	fmt.Println(fact(6))
	fmt.Println(fact(7))
	fmt.Println(fact(8))
}
