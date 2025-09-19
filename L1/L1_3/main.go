package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type Flags struct {
	C int
}

func NewFlags() *Flags {
	return &Flags{}
}

func (f *Flags) Parse() {
	flag.IntVar(&f.C, "c", 3, "set count of workers")
	flag.Parse()
}

type Worker struct {
	ch chan int
	wg sync.WaitGroup
}

func NewWorker() *Worker {
	return &Worker{
		ch: make(chan int, 10),
	}
}

func (w *Worker) Run(n int) {
	for i := 0; i < n; i++ {
		//w.wg.Add(1)
		go func(n, i int) {
			for val := range w.ch {
				time.Sleep(3 * time.Second)
				fmt.Printf("worker %d reade value %d\n", i, val)
			}
		}(n, i)
	}
}

func main() {

	flags := NewFlags()
	flags.Parse()

	workers := NewWorker()
	workers.Run(flags.C)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			var v int
			fmt.Scanf("%d\n", &v)
			workers.ch <- v

		}

	}()
	wg.Wait()
}
