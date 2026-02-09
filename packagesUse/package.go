package main

import (
	"fmt"

	"github.com/Shivang2003/GoLangPractice/packagesUse/pac1"
	"github.com/Shivang2003/GoLangPractice/packagesUse/pac2"
)

func main() {
	pac1.PrintPac1()

	data := pac2.Data{
		Value: 42,
		// isTrue:    true,
		RandomStr: "Hello from package 2",
	}

	fmt.Printf("Data struct from pac2: ", data)

}
