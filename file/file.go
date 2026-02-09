package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	f, err := os.Open("example.txt")

	if(err != nil){
		fmt.Println("Error opening file:", err)
		return
	}


	info, err := f.Stat()

	if(err != nil){
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println(info.Name(), info.Size(), "bytes", info.Mode(), info.IsDir(), info.ModTime())
	
	buff := make([]byte, info.Size())

    d , err := f.Read(buff) //d is buffer length

	if err != nil{
		fmt.Println("Error reading file:", err)
	}
	
	fmt.Println("File content:\n", d, string(buff[:d]))


	dir, err := os.Open("../")

	if err !=  nil{
		fmt.Println("Error opening directory:", err)
		return
	}

	dirInfo , err := dir.Readdir(-1)

	for _, dirItem := range dirInfo{
		fmt.Println("Name ", dirItem.Name())
	}

	var wrr *os.File

	wrr, err = os.Create(fmt.Sprintf("newfile_%d.txt", time.Now().Unix()))
	if err != nil{
		fmt.Println("Error creating file:", err)
		return
	}

	wrr.WriteString("hello qwertyu " + fmt.Sprintf("%d", time.Now().Unix()))



	defer func(){
		f.Close()
		dir.Close()
		wrr.Close()
	} ()

}