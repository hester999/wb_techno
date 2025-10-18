package goexit

import (
	"fmt"
	"runtime"
	"sync"
)

func GoExit() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All Goroutines have completed.")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d is working...\n", id)

	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d is processing task %d\n", id, i)
	}

	runtime.Goexit()
}
