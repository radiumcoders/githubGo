package maps

import "fmt"

func Maps () {
	fmt.Println("maps")
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(m["one"])
	fmt.Println(m["two"])
	fmt.Println(m["three"])
	fmt.Println(len(m))
	delete(m, "one")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
}