package main

import "fmt"


func main(){
i:=1
	for i=1;i<=25; i++ {
		if i%2==0 {
			fmt.Println(i)
		}
	}
	fmt.Println("---------------------------------------------------")
	i=1
	for i<=25 {
		if i%2==0 {
			fmt.Println(i)
		}
		i++;
	}

	fmt.Println("---------------------------------------------------")

	num:=1;
	fmt.Println("Enter a number")
	fmt.Scanln(&num)
	fact:=1;
	for num>0{
		fact= fact* num
		num--
	}
	fmt.Println("factorial of the input number is ", fact)
	fmt.Println("---------------------------------------------------")
	arr:=[5]int{1,2,3,4,5}
	sum:=0
	for _,v:=range arr{
		sum=sum+v
	}
	fmt.Println("The input array is ", arr, "and the sum is ",sum)
	fmt.Println("---------------------------------------------------")
	fmt.Println("original array is  ", arr)
	arr1=make([]int,3)
	fmt.Println("Spliced array is  ", arr)
	arr=append(arr,10,11)
	fmt.Println("Extended array is  ", arr)
	fmt.Println("---------------------------------------------------")
	newmap:=make(map[int]string)
	newmap[6]="Hello"
	newmap[15]="Go"
	newmap[19]="World"
	fmt.Println("Map is  ", newmap)


}


