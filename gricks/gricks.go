package gricks

import "fmt"

type GG[T string | int] struct {
	list []T
}

// T of typr any or we can usr | to loop over any specified data type
// func Loopover[T any](list []T) {
func Loopover[T string | int](list []T) {
	for _, items := range list {
		fmt.Println(items)
	}
}

func G() {
	// can take string or int any of the one
	ex1 := GG[string]{
		list: []string{"apple", "banana", "cherry"},
	}
	// ex2 := GG[int]{
	// 	list: []int{1,2,3,4,5,6,7,8,9,0},
	// }
	for _, item := range ex1.list {
		fmt.Println(item)
	}
	// will give error as bool is not string or int
	// libool := []bool{true, false, true, false, true, false, true, false, true, false}
	// Loopover(libool)
	listring := []string{"apple", "banana", "cherry"}
	linum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	Loopover(linum)
	Loopover(listring)

	fmt.Println(ex1)
}
