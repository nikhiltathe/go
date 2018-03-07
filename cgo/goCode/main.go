package main

// #cgo CFLAGS:  -IcCode/src/include
// #cgo LDFLAGS: -Llib -lbinary
// #include <header.h>
import "C"
import "fmt"

func main() {
	var a C.int
	a := 2
	f := C.Square(a)
	fmt.Println(f)
	// Output: 4
}

// go build --ldflags '-extldflags "-static"'
