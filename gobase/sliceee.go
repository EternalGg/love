package gobase

import "fmt"

func trycopy() {
	a := []int{1, 2, 3, 4}
	b := []int{0}
	c := append(a, 6)
	copy(c, b)
	fmt.Println(c)
}

func sliceadd() {
	a := []int{}
	defer fmt.Println("we done")
	//for i := 0; i < 10; i++ {
	//	a[i] = i
	//}
	fmt.Println(a)

}

func init() {

	//trycopy()
	//sliceadd()
}
