package pointers

import "fmt"

func PointerExample() {
	val := 10
	// when u need to reference memory address of variable we use &
	prt := &val
	// refrence to the memory address of val variable
	fmt.Println(prt)
	// to get the value from the memory address we use *
	fmt.Println(*prt)
	// to declare a pointer variable we use *
	var potr *int
	potr = &val
	fmt.Println(potr)
}