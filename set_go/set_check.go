package main

import (
	"fmt"
	"github.com/fatih/set"
)

func main() {
	// set 是无序不重复的集合，因此每一次存取，读取都是无序随机的

	// 创建一个0 item 的set
	s := set.New(set.ThreadSafe) // thread safe version线程安全
	//s := set.New(set.NonThreadSafe) // non thread-safe version

	// 判定是否为空，是返回true
	s.IsEmpty()

	s.Add(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// 检测单个item是否存在
	s.Has(3)

	// 判定多个item是否全部存在
	s.Has(1, 2, 3, 4)

	t := set.New(set.ThreadSafe)
	t.Add(1, 2, 3, 4)

	// 检测两个set集合是否相同
	if !s.IsEqual(t) {
		fmt.Println("s is not equal to t")
	}

	// 检测t是否是s的子集，也就是是否包含t
	if s.IsSubset(t) {
		fmt.Println("t is a subset of s")
	}

	if t.IsSuperset(s) {
		fmt.Println("s is a superset of t")
	}
}
