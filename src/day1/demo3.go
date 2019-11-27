package main

import "fmt"
//import "reflect"

func main(){
	num1:=0
	num2:=0
	fmt.Println("Enter 2 Numbers ")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Println("Enter Sum of 2 Numbers is ", (num1+num2))
	
}