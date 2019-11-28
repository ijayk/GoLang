package main

import "mypackage"
import "fmt"

func main(){
	ans:=mypackage.Add(1,2)
	fmt.Println("The sum externally calculated is "ans)
}