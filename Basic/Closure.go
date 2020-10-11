package Basic

import "fmt"

func InitSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func TestClosure()  {
	nextInt := InitSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	newInt := InitSeq()
	fmt.Println(newInt())
}