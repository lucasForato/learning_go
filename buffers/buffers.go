package main

import (
	"fmt"
)

type Hello struct {
	foo string
	bar string
}

func (h Hello) String() string {
	return fmt.Sprintf("foo: %v, bar: %v", h.foo, h.bar)
}

func main() {
	hello := &Hello{
		foo: "hey",
		bar: "people",
	}

	fmt.Println(hello)
}
