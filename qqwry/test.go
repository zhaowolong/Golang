package main

import (
	"fmt"

	"code.google.com/p/mahonia"
	"github.com/kayon/qqwry"
)

func main() {
	qw := qqwry.New("./qqwry.dat")
	var res qqwry.Result
	res = qw.Search("114.119.6.83")
	fmt.Printf("IP: %s \nBegin: %s \nEnd: %s \nCountry: %s \nArea: %s \n", res.IP, res.Begin, res.End, res.Country, res.Area)
	enc := mahonia.NewEncoder("UTF-8")
	strr := enc.ConvertString(res.Area)
	fmt.Println(strr)
}
