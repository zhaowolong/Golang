package main

import "fmt"

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

func main() {
	var d data = 15
	var x interface{} = d
	if n, ok := x.(fmt.Stringer); ok {
		fmt.Println(n)
	}

	if n, ok := x.(int); ok {
		fmt.Println(n)
	}
	if d2, ok := x.(data); ok {
		fmt.Println(d2)
	}

	e := x.(error)
	fmt.Println(e)
}
