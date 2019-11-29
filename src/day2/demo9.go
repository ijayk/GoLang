//multithereading in go


package main
import ("fmt";"sync")
var wg=sync.WaitGroup{}

func main(){
	wg.Add(4)
	sayhello()
	saybye()
	fmt.Println("-------------synchronous code complete---------------------")
	go sayhello()
	go saybye()
	wg.Wait()
	 fmt.Println("-------------asynchronous code complete---------------------")
}

func sayhello(){
	i:=0
	for i=0;i<5;i++{
		fmt.Println("sayhello function called", i)
		//time.Sleep(time.Millisecond*1000)
	
	}
	wg.Done()
}

func saybye(){
	i:=0
	for i=0;i<5;i++{
		fmt.Println("saybye function called", i)
		//time.Sleep(time.Millisecond*1000)
	}
	wg.Done()
}

