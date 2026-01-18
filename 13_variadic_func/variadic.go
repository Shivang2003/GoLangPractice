package main
import "fmt"

func sum(nums...int) int {
	s := 0
	for _,i := range(nums){
		s+=i
	}
	return s
}

func main() {
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	nums := []int{11,22,33,44,55,66}
	fmt.Println(sum(nums...))
}