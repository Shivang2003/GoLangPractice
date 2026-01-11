package main
import "fmt"

func main(){
	nums := []int{1,2,3,4,5}

	sum := 0

	for i, num := range(nums){
		sum+=num
		fmt.Println(i, " ", num)
	}
	fmt.Println(sum)

	m := make(map[int]int)
	m[1] = 23
	m[2] = 34
	m[3] = 343

	for k,v := range m {
		println(k, "->",v)
	}

	str := "wererwer"

	for i,c := range str{
		fmt.Println(i,string(c))
	}
}