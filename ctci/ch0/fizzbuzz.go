package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
			fmt.Println()
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	// multiple lines printing
	fmt.Println(`
asdfa
asdf asdf
asdf
`)
}
