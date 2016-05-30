package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	buffer := bytes.NewBufferString("caonima teset heloowdkljsfkljsldjflsdjfldsjflads;fklasdnfo;ewhvoifqwl;j")
	fmt.Println(buffer.String(), buffer.Len())
	fmt.Println(buffer.Bytes())
	bytes := buffer.Bytes()
	bytes[0] = 'c' + 1

	fmt.Println(buffer.String(), buffer.Len())

	//Truncate
	fmt.Println("-----truncate------------------------")
	buffer.Truncate(30)
	fmt.Println(buffer.String(), buffer.Len())

	//Read
	fmt.Println("-----read------------------------")
	var readBuf []byte = make([]byte, 2, 10)
	len, _ := buffer.Read(readBuf)
	fmt.Println(buffer.String(), buffer.Len())
	fmt.Println(readBuf, len)

	buffer.Next(3)
	fmt.Println(buffer.String(), buffer.Len())
	// 吐出最新读出云的一个字节
	buffer.UnreadByte()
	fmt.Println(buffer.String(), buffer.Len())

	// 读取直到包含一个字符的地方
	readStr, _ := buffer.ReadString('h')
	fmt.Println(readStr, readStr)
	fmt.Println(buffer.String(), buffer.Len())

	// 从io读取
	//	buffer.Reset()
	//	buffer.ReadFrom(os.Stdin)
	buffer.WriteTo(os.Stdout)
	var testArray [][2]string
	testArray = make([][2]string, 1)
	testArray[0][0] = "test"
	fmt.Println(testArray[0][0])
}
