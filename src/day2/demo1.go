package main

import "fmt"
import "reflect"

type emp struct {
	id int
	name string
	age float32
}


func main(){

num1:=100
fmt.Println(num1)
//pointer
ptr:=&num1
fmt.Println(ptr)
fmt.Println(reflect.TypeOf(ptr))
*ptr=500
fmt.Println(*ptr)
fmt.Println(num1)
}


