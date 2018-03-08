package main

/*
#cgo CFLAGS: -I${SRCDIR}/headers
#cgo LDFLAGS: ${SRCDIR}/libs/libclibrary.a
#include "clibrary.h"
*/
import "C"

import (
	"fmt"
)

func main() {
	num := 2
	val := C.Square(C.int(num))
	fmt.Println("Square of ", num, " is :", val)
}
