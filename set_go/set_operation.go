package set_go

import (
	"fmt"
	"github.com/fatih/set"
)

func MainOperation() {
	// 初始化a和b
	a := set.New(set.ThreadSafe)
	a.Add("one", "two", "three")
	b := set.New(set.NonThreadSafe)
	b.Add("one", "four", "five")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	// Union并集，也就是a+b
	// ["one", "two", "three", "four", "five"]
	c := set.Union(a, b)
	fmt.Println("a 并集 b = c :", c)

	// 交集
	d := set.Intersection(a, b)
	fmt.Println("a 交集 b = d:", d)

	// 在a里与b不同的items，也就是a-b
	e := set.Difference(a, b)
	fmt.Println("a 差集 b = e:", e)

	// a与b的并集去除a与b的交集，也就是： A∪B - A∩B
	f := set.SymmetricDifference(a, b)
	fmt.Println("A∪B - A∩B = ", f)

	// a与b的交集，并把结果返回给a
	a.Merge(b)
	fmt.Println("a.Merge(b) = ", a)

	// 将数据变为string的slice
	g := set.StringSlice(a)
	fmt.Println("set.StringSlice(a) = ", g)

	// 将数据变为int的slice
	tmp := set.New(set.ThreadSafe)
	tmp.Add(1, 2, 3, 4, 5)
	h := set.IntSlice(tmp)
	fmt.Println("set.IntSlice(a) = ", h)

	// a与b的差集，并把结果返回给a
	a.Separate(b)
	fmt.Println("a.Separate(b) = ", a)
}
