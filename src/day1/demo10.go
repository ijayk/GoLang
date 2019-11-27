package main

import "fmt"


func main(){

	var arr [3]int

	names:=[]string{"David","Letterman","Dravid","Jachuti"}
	arr[0]=10
	arr[1]=20
	arr[2]=30
	fmt.Println("Array has ", arr)

	//Slice

	sli1:=make([]string,2)
	sli1[0]="James"
	sli1[1]="JJ"
	sli2:=append(sli1,"Anil","Sam")
	sli3:=append(sli2,names...)
	fmt.Println(sli3)


	slice1:=make([]string ,2,3)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1[0]="1"
	slice1[1]="2"
	fmt.Println(slice1)

	slice1=append(slice1,"3","4")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1=append(slice1,"5","6")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1=append(slice1,"7","8")
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))



}


