//multithereading in go


package main
import ("fmt";"sync")
var wg=sync.WaitGroup{}

func main(){
	ch:=make(chan int, 5)
	ch<-100
	ch<-500
	ch<-200
	for len(ch)>0 {
		fmt.Println(<-ch)
		fmt.Println(len(ch))
	}
	

}



