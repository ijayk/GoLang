package main

import "fmt"

type emp struct {
	id int
	name string
	age float32
}


func main(){


	e1:=emp{101,"Paul",35.6}
	e2:=emp{102,"Max",40.6}

	//dont have to fill in all
	e3:=emp{id:105,name:"Max"}

	fmt.Println(e1)
	fmt.Println(e2.id,e2.name,e2.age)
	fmt.Println(e3.id,e3.name,e3.age)

}


