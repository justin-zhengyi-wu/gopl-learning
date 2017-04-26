package main

import (
	"flag"
	"fmt"
	//"gopl/ch2/tempconv"
	_ "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	/*
		flag.Parse()
		fmt.Print(strings.Join(flag.Args(), *sep))
		if !*n {
			fmt.Println()
		}
	*/
	//	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	f := 1e100
	i := int(f)
	fmt.Println(i)
}
