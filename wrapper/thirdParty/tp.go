package thirdParty

import "fmt"

type A struct {
}

type A1 struct {
	val int
}

func (a *A) Say() {
	fmt.Println("I am A")
}

func (a *A) Give(v int) *A1 {
	fmt.Println("A Give")
	return &A1{val: v}
}

func (a *A1) SayA1() {
	fmt.Println("I am A1")
	fmt.Printf("%+v\n", a)
	fmt.Println("Value is", a.val)
}
