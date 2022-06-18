package chanel

import "fmt"

func channeltest() {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}

	close(c)
}

func readwrite() {
	c := make(chan int, 5)
	var read chan<- int = c
	var write <-chan int = c
	read <- 2
	fmt.Println(<-write)
}

func init() {
	channeltest()
	//readwrite()
}
