package Panicdefer

import (
	"fmt"
	"sort"
)

//类似的，参照这个创建一个自定义类型的方法，实现这个类型的这三个接口方法，然后在一个这个自定义类型的集合上调用 sort.Sort 方法，我们就可以使用任意的函数来排序 Go 切片了。





// 为了在 Go 中使用自定义函数进行排序，我们需要一个对应的
// 类型。这里我们创建一个为内置 `[]string` 类型的别名的
// `ByLength` 类型，
type ByLength []string
// 我们在类型中实现了 `sort.Interface` 的 `Len`，`Less`
// 和 `Swap` 方法，这样我们就可以使用 `sort` 包的通用
// `Sort` 方法了，`Len` 和 `Swap` 通常在各个类型中都差
// 不多，`Less` 将控制实际的自定义排序逻辑。在我们的例
// 子中，我们想按字符串长度增加的顺序来排序，所以这里
// 使用了 `len(s[i])` 和 `len(s[j])`。
func (s ByLength) Len() int{
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func SortByFunctions()  {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}