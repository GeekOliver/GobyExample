package Basic

import "fmt"

func For()  {
	i := 1

	//while
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//for
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}


	for  {
		fmt.Println("loop")
		break
	}
}
