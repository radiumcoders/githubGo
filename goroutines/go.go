package goroutines

import (
	"fmt"
	"sync"
)

func Print(id int , w *sync.WaitGroup ) {
	defer w.Done()
	fmt.Println("go routines are fun this is => ", id )
}

func Gorutiner() {
	// wait groups :D
	var wg sync.WaitGroup
	for i := range 12 {
		wg.Add(1)
		go Print(i, &wg)
	}
	wg.Wait()
}
