package switchCase

import (
	"fmt"
)

func MultiSwitch() {
	var num int = 3
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd Number")
	case 2, 4, 6, 8, 10:
		fmt.Println("Even Number")
	default:
		fmt.Println("Number not in range 1-10")
	}
}

