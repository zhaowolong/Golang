package main
import "fmt"
func main() {
	mySlice := make([]int, 5, 10)
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("the all arry is", mySlice)
	fmt.Println("len(mySlice):", len(mySlice))
	mySlice = append(mySlice, 1, 2, 3)
	for i, v := range mySlice{
		fmt.Println(i, v)
	}
	fmt.Println("cap(mySlice):", cap(mySlice))

	newSlice := mySlice[:20]
	fmt.Println("newSlice", newSlice)
	slice1 := []int{1, 2, 3, 4, 5}
    slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println("slice2", slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println("slice1", slice1)
}
