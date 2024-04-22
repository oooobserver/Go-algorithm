package concurrent

import (
	"fmt"
	"sync"
)

var (
	mu       sync.Mutex
	put, get = 0, 0
	nums     [100]int
	empty    = make(chan struct{}, 100)
	full     = make(chan struct{}, 100)
)

func producer(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-empty
		mu.Lock()

		nums[put] = i
		fmt.Printf("Put at %d \n", put)
		put++

		full <- struct{}{}
		mu.Unlock()
	}
}

func consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-full
		mu.Lock()

		tmp := nums[get]
		get++
		fmt.Printf("Get at %d\n", tmp)

		empty <- struct{}{}
		mu.Unlock()
	}
}

func Action() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		empty <- struct{}{}
	}

	wg.Add(1)
	wg.Add(1)
	go consumer(&wg)
	go producer(&wg)

	wg.Wait()
}
