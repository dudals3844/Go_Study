package main

import (
	//"fmt"
	//"math"
	//"log"
	//"os"
	"time"
	//"sync"
)

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
			case <- done1:
				println("run1 clear")

			case <- done2:
				println("run2 clear")
				break EXIT
		}
	}
}

func run1(done chan bool){
	time.Sleep(1*time.Second)
	done <- true
}

func run2(done chan bool){
	time.Sleep(2*time.Second)
	done <- true
}
