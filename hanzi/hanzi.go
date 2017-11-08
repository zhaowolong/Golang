package main

import "fmt"
import gpinyin "github.com/jmz331/gpinyin"

func main() {
	const s = "台我要1234!#$翻译成繁体的汉字堡垒asdf"
	r1 := gpinyin.ConvertToPinyinString(s, "-", gpinyin.PINYIN_WITHOUT_TONE)
	var out string
	out = fmt.Sprintf("--------%d", ([]byte(r1))[0])
	fmt.Println(out)
	fmt.Println(r1)
}
