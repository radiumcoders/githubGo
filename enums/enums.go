package enums

import "fmt"

const (
	Pending = iota
	Processing
	Completed
	Failed
)

func Enums() {
	fmt.Println("enums :-", Processing)
}
