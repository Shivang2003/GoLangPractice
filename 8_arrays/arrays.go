package main
import "fmt"

func main(){

	var arr [5]int
	arr[0] = 10

	fmt.Println("Array:", arr[2])

	//array copies
	arrCopy := arr
	arrCopy[2] = 100
	fmt.Println("Original Array:", arr)
	fmt.Println("Copied Array:", arrCopy)

	//array reference
	arrRef := &arr;
	arrRef[2] = 200;
	arrRef2 := arrRef
	arrRef2[3] = 300
	fmt.Println("arrref ",arrRef)
	fmt.Println("arrref2 ",*arrRef2) // * -> actual array value
	fmt.Println("Original arr ", arr)

	

}