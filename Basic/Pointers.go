package Basic

import "fmt"

func zeroval(ival int)  {
	ival = 0
}

func zeroptr(iptr *int)  {
	*iptr = 0
}

func TestPointers()  {
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)//此时的i还是没变，因为调用的函数的i是局部变量，函数结束临时变量i也就没有了

	fmt.Println(&i)

	zeroptr(&i)//解引用
	fmt.Println(i)

	fmt.Println(&i)
}