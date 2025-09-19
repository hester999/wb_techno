package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	count := 0
	for {
		select {
		case <-time.After(1):
			fmt.Println("time out")
			close(ch)
			return
		default:

			count++
			ch <- count
		}
	}

}
