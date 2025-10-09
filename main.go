package main

import (
	"fmt"
	"github.com/radiumcoders/githubGo/funcs"
)

func main() {
	var a float32 = 10.0
	var b float32 = 20.0
	rv := funcs.Multiplication(&a, &b)
	if rv.Error != nil {
		fmt.Println("Error:", rv.Error)
	} else {
		fmt.Println("Result:", rv.Result)
	}
	// num := 10
	// ptr := &num
	// fmt.Println(&ptr)
	// fmt.Println(*ptr)
	// fmt.Println(ptr)
}
