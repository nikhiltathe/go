package main

import "github.com/go/wrapper/wrapper"

func main() {
	obj := wrapper.GetWrapper()
	obj.Say()
}
