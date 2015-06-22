package main

import "fmt"
import (
	"runtime"
	"sync"
)

var a string
var once sync.Once

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func setup() {
	a = "hello, world"
	fmt.Println("once is called")
}
func doprint() {
	once.Do(setup)
	fmt.Println(a)
}
func twoprint() {
	doprint()
	doprint()
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	fmt.Println("end")
	// 计算CPU个数
	fmt.Println("cpu number:", runtime.NumCPU())
	// 测试once
	twoprint()
}
