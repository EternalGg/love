package gorotine

import (
	"fmt"
	"sync"
)

func Goroutine() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("a", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("b", i)
		}
		wg.Done()
	}()

	wg.Wait()
}

func init() {
	Goroutine()
}
