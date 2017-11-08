package main

type data struct{}

func (data) string() string {
	return "data string()"
}

type node struct {
	data interface {
		string() string
	}
}

type stringer interface {
	string() string
}

type tester interface {
	stringer
	test()
}

type teststruct struct {
}

func (self *teststruct) string() string {
	return ""
}

func (self *teststruct) test() {}

type data2 struct {
	x int
}

func main() {
	var t interface {
		string() string
	} = data{}

	n := node{
		data: t,
	}
	println(n.data.string())
	println("----------------------------")
	var d teststruct
	var tt tester = &d
	tt.test()
	tt.string()
	println("----------------------------")
	d2 := data2{1}
	var pd interface{} = d2
	println(pd.(data2).x)
	var pd2 interface{} = &d2
	pd2.(*data2).x = 200
	println(pd.(data2).x)
	println(pd2.(*data2).x)
	println(d2.x)
	println("----------------------------")

}
