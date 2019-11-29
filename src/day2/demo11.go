package main
import ("fmt";"sync")
var wg=sync.WaitGroup{}
func main(){
	wg.Add(2)
ch:=make(chan int)
go read(ch)
go write(ch)

wg.Wait()
}

// func read(c chan int){
	//restricting channel to read
	func read(c <-chan int){

	//fmt.Println(<-c)

cnt:=1
for cnt=1;cnt<6;cnt++{
	fmt.Println("Reader read value",<-c)

}


wg.Done()
}

//func write(c chan int){
	//restricting channel writing
	func write(c chan<- int){
	//time.Sleep(time.Millisecond*10000)
	//c<-100

	cnt:=1
	for cnt=1;cnt<6;cnt++{
		c<-cnt
		fmt.Println("writer sent value",cnt)
	}
	wg.Done()
}