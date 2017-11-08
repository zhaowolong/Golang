package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/yinheli/qqwry"
)

func main() {
	q := qqwry.NewQQwry("qqwry.dat")
	q.Find("218.17.157.4")
	log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)

	q.Find("223.104.63.44")
	log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
	fmt.Println(strings.Contains(q.City, "联通"))
	// output:
	// 2014/02/22 22:10:32 ip:180.89.94.90, Country:北京市, City:长城宽带
}
