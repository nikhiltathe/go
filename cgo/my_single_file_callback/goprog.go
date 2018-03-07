package main

/*
#cgo CFLAGS: -Iheaders
#cgo LDFLAGS: -Llibs -lclibrary

#include "clibrary.h"

*/
import "C"

import (
	"fmt"
)

func main() {
	val := C.Square(C.int(2))
	fmt.Println(val)
}
