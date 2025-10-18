package ctx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func CtxWithCancel() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	stopCh := make(chan os.Signal, 1)
	wg := &sync.WaitGroup{}
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)
	RunCancel(ctx, wg)

	select {
	case <-stopCh:
		cancel()
		wg.Wait()
	}

}

func RunCancel(ctx context.Context, wg *sync.WaitGroup) {

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int, wg *sync.WaitGroup) {
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("call cancel func in Run func for worker %d\n", i)
					return
				case <-ticker.C:
					fmt.Printf("i am worker %d\n", i)
				}
			}
		}(ctx, i, wg)
	}
}
