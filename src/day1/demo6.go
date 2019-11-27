package main

import "fmt"
//import "reflect"

func main(){
	ch:=""
	fmt.Println("Enter a letter")
	fmt.Scanln(&ch)

	switch ch {
	case "a": fmt.Println("It is vowel",ch)
	case "e": fmt.Println("It is vowel",ch)
	case "i": fmt.Println("It is vowel",ch)
	case "o": fmt.Println("It is vowel",ch)
	case "u": fmt.Println("It is vowel",ch)
	default:
		fmt.Println("Not a vowel")
	}
	arr:=[]int{10,3,4,5,6,77,88}
	for i,v:=range arr {
		fmt.Println("Index is",i," Value is",v)
	}
	for _,v:=range arr {
		fmt.Println("Value is",v)
	}
	
}