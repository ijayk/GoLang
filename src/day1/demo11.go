package main

import "fmt"


func main(){

	dict:=make(map[int]string)
	dict[101]="Sam"
	dict[104]="Smith"
	dict[190]="Tracy"
	dict[290]="Jugnu"
	fmt.Println(dict)
	fmt.Println(dict[104])
	dict[104]="Max"
	//
	for k,v:=range dict {
		fmt.Println(k,v)
	}

	delete(dict,104)
	fmt.Println(dict)


}


