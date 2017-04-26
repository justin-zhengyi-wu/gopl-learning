package main

import "fmt"

type Currency int

const (
	USD Currency = 1 + iota // 美元
	EUR                     // 欧元
	GBP                     // 英镑
	RMB                     // 人民币
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}
