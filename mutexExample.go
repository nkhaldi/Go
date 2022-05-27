package main

import (
	"fmt"
	"sync"
)

const N = 1000

func main() {
	// В первом случае значение x меньше будет меньше N,
	// так как горутины перекрывают работу друг друга,
	// и инкремент при одном и том же значении может выполниться дважды.
	fmt.Println("Start without mutex:")
	{
		var x int = 0
		wg := new(sync.WaitGroup)

		for i := 0; i < N; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				x++
			}(wg)
		}

		wg.Wait()
		fmt.Println("Result:", x)
	}

	// В втором случае значение x будет равно N,
	// так как методы Lock и Unlock не позволяют выполнить инкремент дважды.
	fmt.Println("Start with mutex:")
	{
		var x int = 0
		mu := new(sync.Mutex)
		wg := new(sync.WaitGroup)

		for i := 0; i < N; i++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, mu *sync.Mutex) {
				defer wg.Done()
				mu.Lock()
				x++
				mu.Unlock()
			}(wg, mu)
		}

		wg.Wait()
		fmt.Println("Result:", x)
	}
}
