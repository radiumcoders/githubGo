package main

import (
	"fmt"

	// "github.com/radiumcoders/githubGo/funcs"
)

func main() {
	// rv := funcs.Multiplication(12, 0)
	// if rv.Error != nil {
	// 	fmt.Println("Error:", rv.Error)
	// } else {
	// 	fmt.Println("Result:", rv.Result)
	// }
	num := 10 
	ptr := &num
	fmt.Println("Pointer:", ptr)
	fmt.Println("Value:", *ptr)
	fmt.Println(&ptr)
}
