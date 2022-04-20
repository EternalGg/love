package gobase

import (
	"fmt"
	"reflect"
)

func Arithmetic_operators() {
	fmt.Println("lamo")
}

func Var() {
	var str, inT = "string", 0
	fmt.Println("WE ARE", reflect.TypeOf(str), "AND", reflect.TypeOf(inT))
}

func Const() {
	const (
		a = iota
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}

func Array() {
	a := [5]int{0, 1}         // 未初始化元素值为 0。
	b := [...]int{0, 1}       // 通过初始化值确定数组长度
	c := [5]int{2: 10, 4: 20} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(a, b, c, d)
	fmt.Println(len(a), len(b), len(c), len(d))
	fmt.Println(cap(a), cap(b), cap(c), cap(d))
}

func Slice() {
	arr := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-6] //去掉切片的最后一个元素
	slice10 := make([]int, 3, 6)
	fmt.Printf("局部变量： arr2 %v\n", arr)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
	fmt.Println("局部变量： slice10 :", slice10)
}

func Pointer() {
	fmt.Println("-----------------------------------")
	str := "lmao"
	strpointer := &str //指针类型
	strcopy := "lmao"
	fmt.Println(str)
	fmt.Println(&str)
	fmt.Println(strcopy)
	fmt.Println(&strcopy)
	fmt.Println(strpointer)
	fmt.Println(*strpointer)
	str = "lol"
	fmt.Println(str)
	fmt.Println(&str)
	fmt.Println(strpointer)
	fmt.Println(*strpointer)
	fmt.Println("-----------------------------------")
}

func Range() { //range 两个参数分别代表下表和内容
	a := [5]int{1, 2, 3}
	for i, j := range a {
		fmt.Println(a[i], j)
	}

}

//func init() {
//	fmt.Println("we are base")
//	Arithmetic_operators()
//	Var()
//	Const()
//	Array()
//	Range()
//	Slice()
//	Pointer()
//	fmt.Println("base end")
//}
