package main

import (
	"fmt"
	"slices"
)

func main() {

	// ---------------- SLICE CREATION ----------------
	// slice is a dynamic view over an array
	nums := make([]int, 2, 5) // length = 2, capacity = 5

	fmt.Println("Initial slice:")
	fmt.Println("nums:", nums, "len:", len(nums), "cap:", cap(nums))
	fmt.Printf("nums address: %p\n\n", &nums[0])

	// ---------------- APPEND WITHIN CAPACITY ----------------
	nums = append(nums, 23)
	nums = append(nums, 24, 25) // still within capacity

	fmt.Println("After append within capacity:")
	fmt.Println("nums:", nums, "len:", len(nums), "cap:", cap(nums))
	fmt.Printf("nums address: %p\n\n", &nums[0])

	// ---------------- SLICE REFERENCE ----------------
	// numsRef gets a COPY of slice header, not data
	numsRef := nums
	numsRef[0] = 100 // modifies shared array

	fmt.Println("After modifying numsRef[0]:")
	fmt.Println("numsRef:", numsRef)
	fmt.Println("nums   :", nums)
	fmt.Printf("numsRef address: %p\n", &numsRef[0])
	fmt.Printf("nums address   : %p\n\n", &nums[0])

	// ---------------- APPEND AT FULL CAPACITY ----------------
	// capacity is full → append creates NEW array
	numsRef = append(numsRef, 45)

	fmt.Println("After append on numsRef (capacity exceeded):")
	fmt.Println("numsRef:", numsRef)
	fmt.Println("nums   :", nums)
	fmt.Printf("numsRef address: %p\n", &numsRef[0])
	fmt.Printf("nums address   : %p\n\n", &nums[0])

	// ---------------- MODIFY ORIGINAL SLICE ----------------
	nums[1] = 200

	fmt.Println("After modifying nums:")
	fmt.Println("numsRef:", numsRef)
	fmt.Println("nums   :", nums)

	//copy slice
	numsCopy := make([]int, len (nums))
	copy (numsCopy, nums)
	fmt.Println("numsCopy:", numsCopy)
	//compare slices
	fmt.Println(slices.Equal(nums, numsCopy))
}
