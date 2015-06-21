
package main
import "fmt"

func main(){
	str := "hello 世界"
	n := len(str)
	for i:=0; i<n; i++{
		var ch rune  = rune(str[i])
		fmt.Println(i, ch)
	}
	completvalue := 3.2 + 12i
	fmt.Println(real(completvalue), imag(completvalue))
}
