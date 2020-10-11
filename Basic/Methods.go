package Basic

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestMethods()  {
	r := rect{}
	r.width = 10
	r.height = 5

	fmt.Println(r.area())
	fmt.Println(r.perim())


	rp := &r

	fmt.Println(rp.area())
	fmt.Println(r.perim())
}