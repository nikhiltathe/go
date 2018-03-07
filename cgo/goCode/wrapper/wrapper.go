package wrapper

// #cgo CFLAGS:  -I../cCode/src/include
// #cgo LDFLAGS: -Llib -lbinary
// #include <header.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"

func Call() {
	var a C.int
	a := 2
	f := C.Square(a)
	fmt.Println(f)
	// Output: 4
}
