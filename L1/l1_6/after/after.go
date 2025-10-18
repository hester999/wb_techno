package after

import (
	"fmt"
	"sync"
	"time"
)

func After() {

	ch := make(chan int, 1000)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		ch <- i
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			timeAfter := time.After(time.Duration(2) * time.Second)
			defer wg.Done()
			for {
				select {
				case <-timeAfter:
					fmt.Printf("worker %d timeout", i)
					return
				case v := <-ch:
					time.Sleep(1 * time.Second)
					fmt.Printf("worker %d  peocessed %d\n", i, v)
				}
			}

		}(i)
	}
	wg.Wait()
}
