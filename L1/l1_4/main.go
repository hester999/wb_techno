package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ch := make(chan interface{}, 10)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	workerCount := 5

	counter := 0
	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				fmt.Println(v)
			}
		}()
	}

	for {
		select {
		case <-stopCh:
			fmt.Println("stop")
			close(ch)
			wg.Wait()
			return
		default:
			time.Sleep(time.Second * 1)
			counter++
			ch <- counter
		}
	}

}
