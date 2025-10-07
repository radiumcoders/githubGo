package switchCase

import (
	"fmt"
)

func SimpleSwitch() {
	var num int = 5
	switch num {
	case 1: 
		fmt.Println("One")
	case 2: 
		fmt.Println("Two")
	case 3: 
		fmt.Println("Three")
	case 4: 
		fmt.Println("Four")
	case 5: 
		fmt.Println("Five")
	default:
		fmt.Println("Number not in range 1-5")
	}
}