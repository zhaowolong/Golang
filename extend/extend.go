package main
import(
	"fmt"
)

type Dog struct {
	name string	
}

type BDog struct{
	*Dog
	name string
}
func (self *Dog)CallMyName(){
	fmt.Println("the is callMyName is Dob" , self.name)
}

func (self *BDog)CallMyName2(){
	fmt.Println("the is callMyName is bDob" , self.name)
}

type DogC interface{
	CallMyName()
}
func TestCall(dog DogC){
	fmt.Println("-------------------------------")
	dog.CallMyName()
	fmt.Println("-------------------------------")
}
func main(){
	b := &BDog{
		Dog:&Dog{
			name:"this is commom dog",
		},
	}
	TestCall(b)

	dog := &Dog{
		name:"xm",
	}
	TestCall(dog)
}
