package main

import "fmt"

func main() {
	res := SetBit(5, 0, 1)
	fmt.Println(res)
}

func SetBit(n int64, i int, value int) int64 {
	if value == 1 {
		return n | (1 << i)
	}
	return n &^ (1 << i)
}
