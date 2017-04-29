package main

import (
	"fmt"
)

func main() {
	var max int64 = 100
	var a, b int64 = 0, 1
	ret := []int64{0, 1}
	for i := 1; int64(i) < max; i++ {
		t := a + b
		a = b
		b = t
		ret = append(ret, b)
	}
	fmt.Println(ret)
}
