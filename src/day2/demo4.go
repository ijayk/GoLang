package main

import "fmt"

type student struct {
	id int
	name string
	age float32
	saddress []address
}

type address struct{
	pin int
	city string
	state string
}


func main(){
	add1:=address{123,"Mumbai","Maha"}
	add2:=address{1283,"Pune","Maha"}

	add:=[]address{add1,add2}
	sobj:= student{101,"Max",12.4,add}
	fmt.Println(sobj)

	

}