package main

/*
#cgo CFLAGS: -Iheaders
#cgo LDFLAGS: -Llibs -lclibrary

#include "clibrary.h"

*/
import "C"

import (
	"fmt"
	// "unsafe"
)

//export callOnMeGo
func callOnMeGo(in int) int {
	fmt.Printf("Go.callOnMeGo(): called with arg = %d\n", in)
	return in + 1
}

func main() {
	fmt.Printf("Go.main(): calling C function with callback to us\n")
	// C.some_c_func((C.callback_fcn)(unsafe.Pointer(C.callOnMeGo_cgo)))
	val := C.Square(C.int(2))
	fmt.Println(val)
}
