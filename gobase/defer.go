package gobase

import (
	"fmt"
	"os"
)

//https://openhome.cc/Gossip/Go/DeferPanicRecover.html

func FirstDefer(x int) {
	defer fmt.Println("1")
	n := 0
	n++
	var whatever [5]struct{}
	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()
	for i := range whatever {
		defer fmt.Println(i)
	}
	defer fmt.Println(n)
	defer fmt.Println("2")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func SecondDefer() {
	f, err := os.Open("/tmp")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("1")
			fmt.Println(err) // 這已經是頂層的 UI 介面了，想以自己的方式呈現錯誤
		}

		if f != nil {
			if err := f.Close(); err != nil {
				panic(err) // 示範再拋出 panic
			}
		}
	}()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)

	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

func init() {
	SecondDefer()
}
