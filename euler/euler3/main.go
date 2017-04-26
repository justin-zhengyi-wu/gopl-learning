/**
13195的质因数为5,7,13，29.
请问：600851475143 的最大质因数为多少？
*/

package main

import (
	"fmt"
	"math"
)

var primes []int64

func main() {
	const num = 600851475143
	buildPrime(num)
	fmt.Printf("Result: %d\n", primes[len(primes)-1])
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

func buildPrime(n int64) {
	stop := int(math.Floor(math.Sqrt(float64(n))))
	primes = append(primes, 2)
	for i := 3; i < stop; i += 2 {
		isPrime(int64(i))
	}
}
