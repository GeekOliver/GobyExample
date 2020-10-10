package Basic

import (
	"fmt"
)

func Array()  {
	var a [10]int
	fmt.Println(a)

	a[1] = 2
	fmt.Println(a)
	fmt.Println(a[1])

	fmt.Println("lenth:", len(a))
	fmt.Println("cap:", cap(a))


	//
	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	var t [3][6]int
	for i := 0; i < 3; i++{
		for j := 0; j < 6; j++ {
			t[i][j] = i*j + i + j
		}
	}
	fmt.Println(t)
}

func Slice()  {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])

	fmt.Println("len:", len(s))

	s = append(s, "w")
	s = append(s, "wlf")
	fmt.Println(s)


	l := make([]string, len(s))
	copy(l, s)
	fmt.Println(l)

	t := l[1:3]
	fmt.Println(t)

	fmt.Println(s[:5])

	fmt.Println(s[2:])
}

func Map()  {
	m := make(map[string]int)

	m["w"] = 28
	m["z"] = 26

	fmt.Println(m)
	for k, v := range m{
		fmt.Printf("%s: %d\n", k, v)
	}
	v1 := m["w"]
	fmt.Println(v1)

	fmt.Println("len:", len(m))

	delete(m, "z")
	fmt.Println(m)

	n := map[string]int{"1":1, "b":1}
	fmt.Println(n)
}