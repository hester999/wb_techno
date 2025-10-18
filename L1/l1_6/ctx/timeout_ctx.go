package ctx

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func CtxWithTimeout(long int) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(long))
	wg := &sync.WaitGroup{}
	defer cancel()
	RunTimeout(ctx, wg)
	wg.Wait()
}

func RunTimeout(ctx context.Context, wg *sync.WaitGroup) {

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(ctx context.Context, wg *sync.WaitGroup, i int) {
			defer wg.Done()
			ticker := time.NewTicker(time.Second * 1)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					fmt.Println("timeout ")
					return
				case <-ticker.C:
					fmt.Printf("i am worker %d \n", i)
				}
			}
		}(ctx, wg, i)
	}
}
