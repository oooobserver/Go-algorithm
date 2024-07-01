package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				fmt.Println("--------------")
			} else {
				fmt.Println("****************")
			}

		}(i)
	}

	wg.Wait()

}
