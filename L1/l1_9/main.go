package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var arr []int
	for i := 1; i <= 100000; i++ {
		arr = append(arr, i)
	}
	var wg sync.WaitGroup

	wg.Go(func() {
		var threshold int
		for _, v := range arr {
			ch1 <- v
			threshold++

		}
		close(ch1)
	})

	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					close(ch2)
					return
				}
				ch2 <- v * 2
			}
		}
	}()
	go func() {
		wg.Wait()
	}()

	for v := range ch2 {

		fmt.Println(v)
	}

}
