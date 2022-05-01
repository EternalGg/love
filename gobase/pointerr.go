package gobase

import "fmt"

func test1() {
	i1 := test1_1()
	fmt.Print(i1)
}

func test1_1() *int {
	i1 := 3
	fmt.Println(&i1)
	return &i1
}

func test2() {
	i2 := test2_2()
	fmt.Println(&i2)
}

func test2_2() int {
	i2 := 2
	fmt.Println(&i2)
	return i2
}

func init() {
	test2()
}
