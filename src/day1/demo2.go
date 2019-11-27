package main

import "fmt"
import "reflect"
var name string
var age int
var sal float32
func main(){
	name = "Jay"
	age = 25
	sal= 80000.5
	fmt.Println("The name is ", name, "and age is ", age)
	fmt.Println("The age is ", age)
	fmt.Println("The salary is ", sal)
	//local//
	var address string ="Mumbai"
	fmt.Println(address)
	ans:="Correct"
	fmt.Println(ans)
	fmt.Printf("%T","%T","%T",ans,name,age,)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println("Enter name")
	uname:=""
	fmt.Scanln(&uname)
	fmt.Println("The name is ", uname)
}