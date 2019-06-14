package thirdParty

import "fmt"

type A struct {
}

// type A1 struct {
// }

func (a *A) Say() {
	fmt.Println("I am A")
}

// func (a *A) Give() A1 {
// 	return A1{}
// }

// func (a *A1) SayA1() {
// 	fmt.Println("I am A1")
// }
