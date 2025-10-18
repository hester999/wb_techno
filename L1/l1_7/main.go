package main

import (
	"fmt"
	"sync"
)

type GoMap struct {
	m  map[int]string
	mu sync.Mutex
}

func NewGoMap() *GoMap {
	return &GoMap{
		m:  make(map[int]string),
		mu: sync.Mutex{},
	}
}
func main() {

	goMap := NewGoMap()

	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(m *GoMap, i int) {
			defer wg.Done()
			m.mu.Lock()
			defer m.mu.Unlock()
			m.m[i] = fmt.Sprintf("goroutine %d", i)
		}(goMap, i)
	}
	wg.Wait()

	fmt.Println(goMap.m)
}
