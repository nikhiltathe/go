package wrapper

import (
	"github.com/go/wrapper/thirdParty"
)

type B struct {
	*thirdParty.A
}

// func (b *B) Say() {
// 	fmt.Println("in B")
// }

type C struct {
	Wrap Wrap
}

func (b *C) Say() {
	inst = &B{}
	inst.Say()
}

var inst Wrap

func GetWrapper() Wrap {
	return &C{}
}
