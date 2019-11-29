///exception handling in go
package main
import "fmt"


func main(){
	ans:=div(10,0)
	fmt.Println("The division gives ", ans)
	fmt.Println("Last line in main ")
}



func div(n1 int,n2 int)(int){
	defer func(){
		//recover()
		fmt.Println("Recovery from exception done",recover())
	}()
	return n1/n2
}