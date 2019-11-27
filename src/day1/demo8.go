package main

import "fmt"
//import "reflect"

func main(){
sum:= add("This is Go",2,3,4,5,6,7,8,9)
fmt.Println("The sum of the above array is ", sum)
}
func add(str string,arr ...int)int{
	fmt.Println(str)
	fmt.Println(arr)
	sum:=0
	for _,v:=range arr{
		sum = sum +v
	}
	return sum
}