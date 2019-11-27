package main

import "fmt"
//import "reflect"

func main(){
	age:=20
	if age>18 && age<=60 {
		fmt.Println("Person Can Vote")
	}else if age<18 {
		fmt.Println("Person Can Not Vote")
	}else {
		fmt.Println("Invalid Age")
	}
	
}