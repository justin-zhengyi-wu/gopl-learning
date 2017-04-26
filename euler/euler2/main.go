/**
找出 4000000 以下的Fibonacci序数。
如： 1,2,3,5,8,13,21,34,55,89,...
从第三个数字开始，它的值是前两个数字的和，以此类推。
*/

package main

import "fmt"

const MAX2 = 4000000

func main() {
	f1 := 1
	f2 := 2
	sum := 0
	for f2 < MAX2 {
		if isEven(f2) {
			sum += f2
		}
		f1, f2 = f2, f1+f2
	}
	fmt.Printf("Result: %d\n", sum)
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
