package main

import (
	"sync"
	"time"
)

func main() {
	done := make(chan int)
	c := make(chan string)
	go func() {
		s := <-c
		println(s)
		//close(done)
		done <- 1
	}()
	c <- "hi"
	d := <-done
	println(d)
	println("--------------------------------")
	c2 := make(chan int, 3)
	c2 <- 1
	c2 <- 2
	println(<-c2)
	println(<-c2)
	println("--------------------------------")
	var a3, b3 chan int = make(chan int, 3), make(chan int, 3)
	a3 <- 1
	c3 := a3
	println(a3 == b3)
	println(a3 == c3)
	println(c3)
	println(a3)
	println(cap(a3), cap(c3), len(a3))
	println("--------------------------------")
	done4 := make(chan struct{})
	c4 := make(chan int)
	go func() {
		defer close(done4)
		for {
			x, ok := <-c4
			if !ok {
				return
			}
			println(x)
		}
	}()
	c4 <- 1
	c4 <- 2
	c4 <- 3
	//close(c)
	//	<-done4
	println("--------------------------------")
	var wg5 sync.WaitGroup
	ready5 := make(chan struct{})
	for i := 0; i < 3; i++ {
		wg5.Add(1)
		go func(id int) {
			defer wg5.Done()
			println(id, ": ready!")
			<-ready5
			println(id, ":running...")
		}(i)
	}
	println("Ready?Go!")
	time.Sleep(time.Second)
	close(ready5)
	wg5.Wait()
	println("--------------------------------")
	c6 := make(chan int, 3)
	c6 <- 10
	c6 <- 20
	close(c6)
	for i := 0; i < cap(c6)+1; i++ {
		x, ok := <-c6
		println(i, ":", ok, x)
	}

	close(c6)
	println("--------------------------------")

}
