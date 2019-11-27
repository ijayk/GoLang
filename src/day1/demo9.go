package main

import "fmt"


func main(){
	func(){
		fmt.Println("Hello World ",)
	}()
	func(message string){
		fmt.Println("Hello ",message)
	}("Go")
	square:=func(n1 int)int{
		return n1*n1
	}

	display(square(5))

}


func display(ans int){
	fmt.Println(ans)
}
