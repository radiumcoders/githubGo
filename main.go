package main

import (
	"fmt"
	"github.com/radiumcoders/githubGo/variadic"
)

// "github.com/radiumcoders/githubGo/funcs"

func main() {
	a := 12
	defer fmt.Println(variadic.Multiply(&a, &a, &a, &a))
	fmt.Println("Hello, World!")
}
