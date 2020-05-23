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

	// 添加单个item
	s.Add("one")
	s.Add("one") // 添加相同的item是没有任何反应的
	fmt.Println("s.Add() = ", s)

	// 添加多个item
	s.Add("two", "three", "four", 5, 6, 7, 8)
	fmt.Println("s.Add() = ", s)

	// 删除item
	s.Remove("two")
	// 删除不存在的item 是没有反应的
	s.Remove("10")
	fmt.Println("s.Remove() = ", s)

	// 删除多个item
	s.Remove("one", 5, 8)
	fmt.Println("s.Remove() = ", s)

	// 删除一个任意item并且返回item的值
	item := s.Pop()
	fmt.Println("s.Pop() = ", item)

	// 创建一个新的复制对象set继承所有的数值，对应的内存地址不一样
	other := s.Copy()
	fmt.Println("other = ", other)
	fmt.Println("s = ", s)
	// 集合的串表示
	fmt.Printf("set is %s\n", s.String())

	// remove all items
	s.Clear()
	fmt.Println("s.Clear() = ", s)

	// number of items in the set
	len := s.Size()
	fmt.Println("s len = ", len)

	// return a list of items
	itemList := s.List()
	for item := range itemList {
		fmt.Printf("%v ", item)
	}
	fmt.Println()

}
