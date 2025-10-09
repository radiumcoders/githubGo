package main

import (
	"fmt"

	"github.com/radiumcoders/githubGo/funcs"
)

func main() {
	rv := funcs.Multiplication(12, 0)
	if rv.Error != nil {
		fmt.Println("Error:", rv.Error)
	} else {
		fmt.Println("Result:", rv.Result)
	}
}
