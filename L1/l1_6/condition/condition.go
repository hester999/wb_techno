package condition

import (
	"fmt"
	"sync"
)

func Condition() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- -1
	ch <- 3
	ch <- 4

	close(ch)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			if v < 0 {
				fmt.Println("invalid num")
				return
			}
			fmt.Println("process num:", v)
		}
	}()
	wg.Wait()
}
