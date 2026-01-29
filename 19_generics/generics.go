package main 

import "fmt"

func printInfo[T any](items []T){
	for _,item :=range items{
		fmt.Println(item)	
	}
}

//for scoping generics to main function

func printInfoTwo[T int | float64 | string](items []T){
	for _,item :=range items{
		fmt.Println(item)	
	}
}

type Stack[T int | string] struct{
	elements []T
}

func main(){
	numbs := []int{1,2,3,5}
	printInfo(numbs)

	strs := []string{"go","is","awesome"}
	printInfo(strs)

	newStack := Stack[int]{
		elements: []int{10,20,30},
	}
	fmt.Println("Stack elements:", newStack.elements)
}