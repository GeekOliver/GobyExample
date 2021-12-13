package Basic

import (
	"fmt"
	"math"
)

func Value()  {
	fmt.Println("go" + "lang")

	//整数 float
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	//bool
	fmt.Println(true && false)
	fmt.Println(true||false)
	fmt.Println(!true)
}

func Variables()  {
	var a string = "init"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//默认初始化为零值
	var e int
	fmt.Println(e)

	//简写
	f := "short"
	fmt.Println(f)
}

const s string = "constant"

func Constant()  {
	fmt.Println(s)

	const n = 500000000

	const d =3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}