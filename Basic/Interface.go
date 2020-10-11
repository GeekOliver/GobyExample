package Basic

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2*s.width * 2*s.height
}


// `circle` 的实现。
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量的是接口类型，那么我们可以调用这个被命名的
// 接口中的方法。这里有一个一通用的 `measure` 函数，利用这个
// 特性，它可以用在任何 `geometry` 上。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterface() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	// 结构体类型 `circle` 和 `square` 都实现了 `geometry`
	// 接口，所以我们可以使用它们的实例作为 `measure` 的参数。
	measure(s)
	measure(c)
}
