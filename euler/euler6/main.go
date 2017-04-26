/**
找出200万以下的所有质数的总和。
结果：142913828922
*/

package main

import (
	"fmt"
)

var primes []int64

func main() {
	const MAX = 2000000
	sum := 2
	primes = append(primes, 2)
	// 跳过所有偶数
	for i := 3; i < MAX; i += 2 {
		if isPrime(int64(i)) {
			sum += i
		}
	}
	fmt.Println("Result: ", sum)
}

func isPrime(n int64) bool {
	var stop = len(primes)
	for i := 0; i < stop; i++ {
		if n%primes[i] == 0 {
			return false
		}
		if primes[i]*primes[i] > n {
			break
		}
	}
	primes = append(primes, n)
	return true
}
