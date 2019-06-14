package wrapper

import (
	"fmt"

	"github.com/go/wrapper/thirdParty"
)

var inst Wrap

type B struct {
	*thirdParty.A
}

type B1 struct {
	*thirdParty.A1
}

// func (b *B) Say() {
// 	fmt.Println("in B")
// }

func (b *B) Give(v int) Wrap2 {
	fmt.Println("Giving by B")
	return &B1{A1: (b.A.Give(v))}
}

func (b *B1) SayA1() {
	fmt.Println("B1 SayA1")
	b.A1.SayA1()
}

func GetWrapper() Wrap {
	inst = &B{}
	return inst
}
