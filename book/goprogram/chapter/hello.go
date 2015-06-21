package main
import "fmt"

func GetName()(f, m, l string){
	return "z", "w", "l"

}
func main(){
	var i int = 2
	j := 3
	fmt.Println("hello world", i, j)
	var zhao, wo, long = GetName()
	fmt.Println(zhao, wo, long)
	var _, _, m = GetName()
	fmt.Println(m)
}

