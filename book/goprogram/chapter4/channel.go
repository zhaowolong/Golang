package main

import (
	"fmt"
)

func Count(ch chan int) {
	fmt.Println("countting begin")
	ch <- 2
	fmt.Println("countting end")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		fmt.Println("read begin")
		i := <-ch
		fmt.Println("read end", i)
	}
}
