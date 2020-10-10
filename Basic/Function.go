package Basic

import "fmt"

func Vals() (int, int) {
	return 1,2
}

//Basic.Sum(10, 20, 30, 50)
//
//nums := []int{1,2,3,4}
//Basic.Sum(nums...)
func Sum(nums ...int)  {
	fmt.Print(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)
}