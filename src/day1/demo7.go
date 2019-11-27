package main

import "fmt"

func main(){
add(10,20)
ans:= mult(20,50)
fmt.Println("The result is ", ans)
ans1,ans2:=addmult(12,92)
fmt.Println("The addition is ", ans1)
fmt.Println("The multiplication is ", ans2)
ans3, ans4:= divide(210,0)
if ans4==""{
	fmt.Println("The multiplication is ", ans3)
}else{
fmt.Println("There was an error in the operation")
}
}

func add(num1 int, num2 int){
	fmt.Println("The sum is ", (num1+num2))
}
func mult(num1 int, num2 int)int{
	return num1*num2
}
//multiple args return
func addmult(num1 int, num2 int)(int, int){
	return num1+num2,num1*num2
}

func divide(num1 int, num2 int)(int, string){
	message:=""
	if num2==0{
		message="Error"
		return 0, message
	}
	return (num1/num2), message
}