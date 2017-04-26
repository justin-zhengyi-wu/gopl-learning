/**
找出 1000 以下的 3 或 5 的倍数的数字，并计算出这些数字的总和。
结果: 233168
*/

package main

import (
	"fmt"
)

const MAX = 1000

func main() {
	sum := 0
	for i := 1; i < MAX; i++ {
		sum += check(i)
	}
	fmt.Printf("The sum is %d\n", sum)
}

func check(i int) int {
	if (i%3 == 0) || (i%5 == 0) {
		return i
	}
	return 0
}
