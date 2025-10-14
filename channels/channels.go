package channels

import (
	"fmt"
	"math/rand"
	"sync"
)

func HelprCFunc(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Helper function called", <- numChan)
}

func AddChan(numChan chan int, num1 int , num2 int , wg *sync.WaitGroup){
	result := num1 + num2
	numChan <- result
	wg.Done()
}

func ChannelExample(){
	var wg sync.WaitGroup
	numChan := make(chan int)
	wg.Add(1)
	num1 := rand.Intn(100)
	num2 := rand.Intn(100)
	go AddChan(numChan, num1, num2, &wg)
	// go HelprCFunc(numChan, &wg)
	// numChan <- 435
	res := <- numChan
	fmt.Println(res)
	wg.Wait()
}