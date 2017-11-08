package main

import "fmt"

type sTest struct {
	x int
	s string
}

func test(x *int) {
	fmt.Printf("Pointer:%p, target :%v, %d\n", &x, x, *x)
}

func main() {

	a := sTest{x: 1}
	fmt.Printf("%T, %v", a, a)
	var (
		i int
		s string
		b bool
	)

	ii, ss, bb := 0, "", false
	println(i, s, b)
	println(ii, ss, bb)
	x, y := 1, 2
	x, y = y+3, x+2
	println(x, y)

	const cx, cy = 100, 200
	const (
		cxx uint16 = 12
		cyy
		css = "abc"
		czz
	)

	fmt.Println("%T, %v", cyy, cyy)
	println("")
	fmt.Println("%T, %v", czz, czz)

	af := 0x100
	p := &af
	fmt.Printf("Pointer:%p, target :%v, %d\n", &p, p, *p)
	test(p)
}
