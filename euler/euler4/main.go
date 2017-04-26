/**

 */
package main

import "fmt"

var primes []int64

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

func main() {
	const n = 10001
	const MAX = 1000000
	c := 1
	for i := 3; i < MAX; i += 2 {
		if isPrime(int64(i)) {
			c += 1
		}
		if c == n {
			fmt.Printf(" QED: %v \n", i)
			fmt.Println("--")
			break
		}
	}
}
