package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	var index int = 1
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		case ch <- 3:
		case ch <- 4:
		}
		i := <-ch
		fmt.Println("Value received:", i)
		index = index + 1
		if index > 100 {
			break
		}
	}
}
