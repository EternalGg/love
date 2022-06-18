package closesuree

import "fmt"

func closesure() {
	kid()()
	call := kid()
	call()
	add := add(1, 2)
	add()
}

func kid() func() {
	return func() {
		fmt.Println("no i dont")
	}
}

func add(a int, b int) func() {
	return func() {
		fmt.Println(a + b)
	}
}

func init() {
	//closesure()
}
