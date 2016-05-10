package main
import(
	"fmt"
	"strconv"
)

func main(){
	// test parseint
	nNbr, _ := strconv.ParseInt("132342", 0, 64)
	fmt.Println(nNbr)
}
