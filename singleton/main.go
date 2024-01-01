package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance(i int) *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Printf("Creating single instance now. goroutine number: %d \n", i)
			singleInstance = &single{}
		} else {
			fmt.Printf("Single instance already created 1. goroutine number: %d \n", i)
		}
	} else {
		fmt.Printf("Single instance already created 2. goroutine number: %d \n", i)
	}

	return singleInstance
}

// Notice that i passed the itearation number as parameter to the function getInstance(i)
// It was just for sake of curiosity, by doing this we can check the number of goroutine
// that is calling the method and realize how is the assyncronism
func main() {

	for i := 0; i < 30; i++ {
		go getInstance(i)
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
