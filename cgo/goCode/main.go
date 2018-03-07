package main

import (
	"github.com/go/cgo/cgoexample"
	// "github.com/go/cgo/goCode/wrapper"
)

func main() {
	// wrapper.Call()
	cgoexample.Example()
}

// go build --ldflags '-extldflags "-static"'
