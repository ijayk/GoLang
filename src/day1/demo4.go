package main

import "fmt"
//import "reflect"

func main(){
	cnt:=0
	//for loop
	for cnt=0;cnt<10;cnt++ {
		fmt.Println(cnt)
	}
	///while with for
	i:=10
	for i=10;i>0;i-- {
		fmt.Println("From while for ",i)
		//i--
	}
	
}