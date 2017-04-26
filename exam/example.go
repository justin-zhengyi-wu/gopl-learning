/**
集合相关的方法
*/

package main

import (
	"fmt"
)

func main() {

}

// 返回第一个查找到的数组中的元素的位置，如无返回-1
func Index(c []string, el string) int {
	for i, v := range c {
		if el == v {
			return i
		}
	}
	return -1
}

// 集合c中是否含有元素el
func Include(c []string, el string) bool {
	return Index(c, el) >= 0
}

// 集合中至少有一个元素满足第二个参数传入的函数条件
func Any(c []string, f func(string) bool) bool {
	for _, v := range c {
		if f(v) {
			return true
		}
	}
	return false
}

//
func All(c []string, f func(string) bool) bool {
	for _, v := range c {
		if !f(v) {
			return false
		}
	}
	return true
}

// 自定义方法来做过滤
func Filter(c []string, f func(string) bool) []string {
	ret := make([]string, len(c))
	for _, v := range c {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

//
func Map(c []string, f func(string) string) []string {
	ret := make([]string, len(c))
	for _, v := range c {
		ret[i] = f(v)
	}
	return ret
}
