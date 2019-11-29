///defer keyword in go
//if all are defered then LIFO
package main
import "fmt"


func main(){
	defer add()
	sub()
	div()
}

func sub(){
	fmt.Println("Substract funtion called")
}

func div(){
	fmt.Println("Divide funtion called")
}


func add(){
	fmt.Println("Add funtion called")
}