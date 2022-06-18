package _interface

import "fmt"

type humen interface {
	play()
	eat()
	sleep()
}

type babe struct {
	name string
	age  int
}

func (bb babe) eat() {
	fmt.Println(bb.name, "在喝奶")
}
func (bb babe) play() {
	fmt.Println(bb.name, "在wan")
}
func (bb babe) sleep() {
	fmt.Println(bb.name, "在sj")
}

func interfacee() {
	var god humen
	god = babe{
		"mg",
		100,
	}
	god.sleep()
	god.play()
	god.eat()
}

func init() {

	//interfacee()
}
