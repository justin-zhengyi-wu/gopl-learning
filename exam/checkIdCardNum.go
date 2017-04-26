package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input your ID card number")
		return
	}
	checkIdCard(os.Args[1])
}

func checkIdCard(a string) {
	if len(a) != 18 {
		fmt.Println("Invalid ID length")
		return
	}
	a = strings.ToUpper(a)
	sum := 0
	x := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	m := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	for i := 0; i < len(a)-1; i++ {
		el, _ := strconv.Atoi(a[i : i+1])
		el = el * x[i]
		sum += el
	}
	if a[17:] == m[(sum%11)] {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}

}
