package main 
import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("hello world") 
	var weektest time.Weekday
	weektest = time.Monday
	fmt.Println("hello world %s", weektest.String()) 
	loc, _ := time.LoadLocation("USA")
	fmt.Println("location", loc.String())
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	t = time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}
	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
	t0 := time.Now()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

}
