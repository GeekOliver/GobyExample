package Basic

import "fmt"

func IfElse(n int)  {
	if n%2 == 0 {
		fmt.Printf("%d is even", n)
	}else {
		fmt.Printf("%d is odd", n)
	}

	if n % 4 == 0 {
		fmt.Printf("%d is divisible by 4", n)
	}

	if n < 0 {
		fmt.Printf("%d  is negative", n)
	}else if n < 10 {
		fmt.Printf("%d  has 1 digit", n)
	}else {
		fmt.Println(n, " has multiple digits")
	}
}