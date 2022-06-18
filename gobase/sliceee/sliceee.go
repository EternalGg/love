package sliceee

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

func testslicerange() {
	s := make([]int, 3)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println(&s)
	fmt.Println(&s[0], &s[1], &s[2])
	for i, v := range s {
		s[i] = 8

		fmt.Println(&v, v, &s[i], s[i])
		fmt.Println(s[0])
	}

	fmt.Println(s)
	return
}

func init() {
	//testslicerange()
	trycopy()
	//sliceadd()
}
