package main
import "fmt"

func main() {
	//while
	i := 0

	for i<=3 {
		fmt.Println("while loop i:", i)
		i++
	}

	//classic for
	for j :=0; j<3; j++ {
		fmt.Println("classic for loop j:", j)
	}

	//range
	for x := range 3 {
		fmt.Println("range loop x:", x)
	}

	nums := []int{10,20,30,40}
	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}

	keyVals := map[int]string{1:"one", 2:"two", 3:"three"}
	for key , val := range keyVals{
		fmt.Println("key ",key," ","value ",val)
	}
}