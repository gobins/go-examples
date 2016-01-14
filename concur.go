package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//To declare the max number of processes available to our code
	runtime.GOMAXPROCS(2)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("I'm the other function")
	}()

	waitGrp.Wait()
}
