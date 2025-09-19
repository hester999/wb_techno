package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}
	ich := make(chan int, len(arr))
	och := make(chan int, len(arr))

	for _, v := range arr {
		ich <- v
	}
	close(ich)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ich {
				och <- v * v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(och)
	}()

	for v := range och {
		fmt.Println(v)
	}

}
