package main

import (
	"fmt"
	"sync"
)

type N int

func (n N) Value() {
	n++
	fmt.Printf("v:%p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("v:%p, %v\n", n, *n)
}

//--------------------------------
type X struct {
	a int
}

func (self X) test() {
	fmt.Println("hello", self, self.a)
}

//--------------------------------
type data struct {
	sync.Mutex
	buf [1024]byte
}

//--------------------------------
type user struct {
}

type manager struct {
	user
}

func (user) tostring() string {
	return "user"
}

func (m manager) tostring() string {
	return m.user.tostring() + "  manager"
}

//--------------------------------
func main() {

	var a N = 25
	(&a).Value()
	a.pointer()
	a.pointer()
	fmt.Printf("a: %p, %v\n", &a, a)

	fmt.Println("---------------------")
	var xa *X = &X{}
	xb := X{}
	xa.test()
	xb.test()
	(X{}).test()
	fmt.Println("---------------------")

	fmt.Println("---------------------")
	d := data{}
	d.Lock()
	defer d.Unlock()
	fmt.Println("---------------------")

	fmt.Println("---------------------")
	var m manager
	println(m.tostring())
	println(m.user.tostring())
	fmt.Println("---------------------")
}
