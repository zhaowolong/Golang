package main

import "fmt"
import "reflect"

type X int
type strcutTest struct {
	x X
	y X
}

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var a X = 100
	var b strcutTest
	t := reflect.TypeOf(a)
	ts := reflect.TypeOf(b)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println(ts.Name(), ts.Kind())
	fmt.Println(t.String())
	fmt.Println(ts.String())
	fmt.Println(t.Align())
	fmt.Println(ts.Align())

	arrayA := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	mapM := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Println(arrayA, mapM)

	fmt.Println("--------------------------")
	x := 100
	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println(tx, tp, tx == tp)
	fmt.Println(tx.Kind(), tp.Kind())
	fmt.Println(tx.Kind(), tp.Elem())
	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())
	fmt.Println("--------------------------")

	var m3 manager
	t3 := reflect.TypeOf(&m3)
	if t3.Kind() == reflect.Ptr {
		t3 = t3.Elem()
	}
	for i := 0; i < t3.NumField(); i++ {
		f := t3.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)
		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println("   ", af.Name, af.Type)
			}
		}
	}
	fmt.Println("--------------------------")
}
