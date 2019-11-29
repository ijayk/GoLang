//check package runime(runtime environment) time(date time) string(string manipulation) 

package main

import (
	"io/ioutil"
	"fmt"
	"os"
	
)
func main(){
	err:= os.Mkdir("myFolder",0644)
	fmt.Println("Folder Created")
	fp, err := os.Create("MyFile1.txt")
	if err!= nil {
		fmt.Println("File created", fp)
	}
	fp.WriteString("This is a line")
	fp.WriteString("\nThis is a line")
	fp.WriteString("\nThis is a line")
	data:=[]byte("\nMy name is Sam \n and i am here ")
	fp.Write(data)
	fmt.Println("File written to")
	fp.Close()

	///read file
	data11,_:=ioutil.ReadFile("MyFile1.txt")
	fmt.Println("Data from file",string(data11))


}