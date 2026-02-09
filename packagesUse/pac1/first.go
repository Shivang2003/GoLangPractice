package pac1

import "fmt"

func provideString() string {
 return "Hello from package 1"
}

func PrintPac1(){
	fmt.Println(provideString())
}