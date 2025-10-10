package main

import (
	maps "github.com/radiumcoders/githubGo/maps_loops"
	"github.com/radiumcoders/githubGo/pointers"
)

// "github.com/radiumcoders/githubGo/funcs"

func main() {
	// var a float32 = 10.0
	// var b float32 = 0
	// rv := funcs.Multiplication(&a, &b)
	// if rv.Error != nil {
	// 	fmt.Println("Error:", rv.Error)
	// } else {
	// 	fmt.Println("Result:", rv.Result)
	// }
	// a := 12
	// fmt.Println(a)
	// // converting int to string don't give error
	// b := strconv.Itoa(a)
	// fmt.Println("Result:", b)
	// // converting string to int
	// c , err := strconv.Atoi(b)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(c)
	pointers.PointerExample()
	maps.Maps()
}
